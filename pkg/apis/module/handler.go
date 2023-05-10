package module

import (
	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/pkg/apis/module/view"
	"github.com/seal-io/seal/pkg/apis/runtime"
	modbus "github.com/seal-io/seal/pkg/bus/module"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/module"
	"github.com/seal-io/seal/pkg/dao/types/status"
	"github.com/seal-io/seal/pkg/topic/datamessage"
	"github.com/seal-io/seal/utils/topic"
)

func Handle(mc model.ClientSet) Handler {
	return Handler{
		modelClient: mc,
	}
}

type Handler struct {
	modelClient model.ClientSet
}

func (h Handler) Kind() string {
	return "Module"
}

func (h Handler) Validating() any {
	return h.modelClient
}

// Basic APIs.

func (h Handler) Create(ctx *gin.Context, req view.CreateRequest) (view.CreateResponse, error) {
	entity := req.Model()

	creates, err := dao.ModuleCreates(h.modelClient, entity)
	if err != nil {
		return nil, err
	}
	entity, err = creates[0].Save(ctx)
	if err != nil {
		return nil, err
	}

	if err = modbus.Notify(ctx, entity); err != nil {
		return nil, err
	}

	return model.ExposeModule(entity), nil
}

func (h Handler) Delete(ctx *gin.Context, req view.DeleteRequest) error {
	return h.modelClient.Modules().DeleteOne(req.Model()).Exec(ctx)
}

func (h Handler) Update(ctx *gin.Context, req view.UpdateRequest) error {
	prev, err := h.modelClient.Modules().Get(ctx, req.ID)
	if err != nil {
		return err
	}

	var (
		entity = req.Model()
		// Sync schema on source/version updates.
		shouldSyncSchema = prev.Source != entity.Source
	)

	if shouldSyncSchema {
		entity.Status = status.ModuleStatusInitializing
		entity.StatusMessage = ""
	}

	update, err := dao.ModuleUpdate(h.modelClient, entity)
	if err != nil {
		return err
	}
	if _, err = update.Save(ctx); err != nil {
		return err
	}

	if !shouldSyncSchema {
		return nil
	}

	return modbus.Notify(ctx, entity)
}

func (h Handler) Get(ctx *gin.Context, req view.GetRequest) (view.GetResponse, error) {
	entity, err := h.modelClient.Modules().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return model.ExposeModule(entity), nil
}

// Batch APIs.

func (h Handler) CollectionDelete(ctx *gin.Context, req view.CollectionDeleteRequest) error {
	return h.modelClient.WithTx(ctx, func(tx *model.Tx) (err error) {
		for i := range req {
			err = tx.Modules().DeleteOne(req[i].Model()).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
		return
	})
}

var (
	queryFields = []string{
		module.FieldID,
	}
	getFields  = module.Columns
	sortFields = []string{
		module.FieldID,
		module.FieldStatus,
		module.FieldCreateTime,
	}
)

func (h Handler) CollectionGet(
	ctx *gin.Context,
	req view.CollectionGetRequest,
) (view.CollectionGetResponse, int, error) {
	query := h.modelClient.Modules().Query()
	if queries, ok := req.Querying(queryFields); ok {
		query.Where(queries)
	}

	// Get count.
	cnt, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get entities.
	if limit, offset, ok := req.Paging(); ok {
		query.Limit(limit).Offset(offset)
	}
	if fields, ok := req.Extracting(getFields, getFields...); ok {
		query.Select(fields...)
	}
	if orders, ok := req.Sorting(sortFields, model.Desc(module.FieldCreateTime)); ok {
		query.Order(orders...)
	}
	entities, err := query.
		// Allow returning without sorting keys.
		Unique(false).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeModules(entities), cnt, nil
}

func (h Handler) CollectionStream(ctx runtime.RequestUnidiStream, req view.CollectionStreamRequest) error {
	t, err := topic.Subscribe(datamessage.Module)
	if err != nil {
		return err
	}
	defer func() { t.Unsubscribe() }()

	query := h.modelClient.Modules().Query()
	if fields, ok := req.Extracting(getFields, getFields...); ok {
		query.Select(fields...)
	}

	for {
		var event topic.Event
		event, err = t.Receive(ctx)
		if err != nil {
			return err
		}
		dm, ok := event.Data.(datamessage.Message[string])
		if !ok {
			continue
		}

		var streamData view.StreamResponse
		switch dm.Type {
		case datamessage.EventCreate, datamessage.EventUpdate:
			entities, err := query.Clone().
				// Allow returning without sorting keys.
				Unique(false).
				Where(module.IDIn(dm.Data...)).
				All(ctx)
			if err != nil {
				return err
			}
			streamData = view.StreamResponse{
				Type:       dm.Type,
				Collection: model.ExposeModules(entities),
			}
		case datamessage.EventDelete:
			streamData = view.StreamResponse{
				Type: dm.Type,
				IDs:  dm.Data,
			}
		}
		if len(streamData.IDs) == 0 && len(streamData.Collection) == 0 {
			continue
		}
		err = ctx.SendJSON(streamData)
		if err != nil {
			return err
		}
	}
}

// Extensional APIs.

func (h Handler) RouteRefresh(ctx *gin.Context, req view.RefreshRequest) error {
	m, err := h.modelClient.Modules().Get(ctx, req.ID)
	if err != nil {
		return err
	}
	m.Status = status.ModuleStatusInitializing
	m.StatusMessage = ""
	update, err := dao.ModuleUpdate(h.modelClient, m)
	if err != nil {
		return err
	}
	if err = update.Exec(ctx); err != nil {
		return err
	}

	return modbus.Notify(ctx, m)
}
