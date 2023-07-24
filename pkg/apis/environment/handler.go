package environment

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/pkg/apis/environment/view"
	envbus "github.com/seal-io/seal/pkg/bus/environment"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/project"
	"github.com/seal-io/seal/pkg/dao/types/object"
	pkgservice "github.com/seal-io/seal/pkg/service"
	"github.com/seal-io/seal/utils/log"
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
	return "Environment"
}

func (h Handler) Validating() any {
	return h.modelClient
}

// Basic APIs.

func (h Handler) Create(ctx *gin.Context, req view.CreateRequest) (view.CreateResponse, error) {
	entity := req.Model()

	err := h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		creates, err := dao.EnvironmentCreates(tx, entity)
		if err != nil {
			return err
		}

		entity, err = creates[0].Save(ctx)
		if err != nil {
			return err
		}

		serviceInputs := make(model.Services, 0, len(req.Services))

		for _, s := range req.Services {
			svc := s.Model()
			svc.EnvironmentID = entity.ID
			serviceInputs = append(serviceInputs, svc)
		}

		if err := pkgservice.SetSubjectID(ctx, serviceInputs...); err != nil {
			return err
		}

		services, err := pkgservice.CreateScheduledServices(ctx, tx, serviceInputs)
		if err != nil {
			return err
		}

		entity.Edges.Services = services

		return envbus.NotifyIDs(ctx, tx, envbus.EventCreate, entity.ID)
	})
	if err != nil {
		return nil, err
	}

	return model.ExposeEnvironment(entity), nil
}

func (h Handler) Delete(ctx *gin.Context, req view.DeleteRequest) error {
	return h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		env, err := dao.GetEnvironmentByID(ctx, tx, req.ID)
		if err != nil {
			return err
		}

		if err = tx.Environments().DeleteOne(req.Model()).Exec(ctx); err != nil {
			return err
		}

		if err = envbus.Notify(ctx, tx, envbus.EventDelete, model.Environments{env}); err != nil {
			// Proceed on clean up failure.
			log.Warnf("environment post deletion hook failed: %v", err)
		}

		return nil
	})
}

func (h Handler) Update(ctx *gin.Context, req view.UpdateRequest) error {
	entity := req.Model()

	return h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		updates, err := dao.EnvironmentUpdates(tx, entity)
		if err != nil {
			return err
		}

		if err = updates[0].Exec(ctx); err != nil {
			return err
		}

		return envbus.NotifyIDs(ctx, tx, envbus.EventUpdate, req.ID)
	})
}

func (h Handler) Get(ctx *gin.Context, req view.GetRequest) (view.GetResponse, error) {
	entity, err := h.modelClient.Environments().Query().
		Where(environment.ID(req.ID)).
		WithProject(func(pq *model.ProjectQuery) {
			pq.Select(project.FieldName)
		}).
		WithConnectors(func(rq *model.EnvironmentConnectorRelationshipQuery) {
			rq.Order(model.Desc(environmentconnectorrelationship.FieldCreateTime)).
				Select(environmentconnectorrelationship.FieldEnvironmentID).
				// Allow returning without sorting keys.
				Unique(false).
				Select(environmentconnectorrelationship.FieldConnectorID).
				WithConnector(
					func(cq *model.ConnectorQuery) {
						cq.Select(
							connector.FieldID,
							connector.FieldType,
							connector.FieldName)
					})
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return model.ExposeEnvironment(entity), nil
}

// Batch APIs.

func (h Handler) CollectionDelete(ctx *gin.Context, req view.CollectionDeleteRequest) error {
	return h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		ids := make([]object.ID, len(req))
		for i := range req {
			ids[i] = req[i].ID
		}

		envs, err := dao.GetEnvironmentsByIDs(ctx, tx, ids...)
		if err != nil {
			return err
		}

		if len(envs) != len(req) {
			return errors.New("invalid id: data mismatch")
		}

		if _, err = tx.Environments().Delete().
			Where(environment.IDIn(ids...)).
			Exec(ctx); err != nil {
			return err
		}

		if err = envbus.Notify(ctx, tx, envbus.EventDelete, envs); err != nil {
			// Proceed on clean up failure.
			log.Warnf("environment post deletion hook failed: %v", err)
		}

		return nil
	})
}

var (
	queryFields = []string{
		environment.FieldName,
	}
	getFields = environment.WithoutFields(
		environment.FieldUpdateTime)
	sortFields = []string{
		environment.FieldName,
		environment.FieldCreateTime,
	}
)

func (h Handler) CollectionGet(
	ctx *gin.Context,
	req view.CollectionGetRequest,
) (view.CollectionGetResponse, int, error) {
	query := h.modelClient.Environments().Query()
	if queries, ok := req.Querying(queryFields); ok {
		query.Where(queries)
	}

	if req.ProjectID != "" {
		query.Where(environment.ProjectID(req.ProjectID))
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

	if orders, ok := req.Sorting(
		sortFields,
		model.Desc(environment.FieldCreateTime),
	); ok {
		query.Order(orders...)
	}

	entities, err := query.
		// Allow returning without sorting keys.
		Unique(false).
		// Must extract connectors.
		Select(environment.FieldID).
		WithConnectors(func(rq *model.EnvironmentConnectorRelationshipQuery) {
			// Includes connectors.
			rq.Order(model.Desc(environmentconnectorrelationship.FieldCreateTime)).
				WithConnector(func(cq *model.ConnectorQuery) {
					cq.Select(
						connector.FieldID,
						connector.FieldType,
						connector.FieldName)
				})
		}).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeEnvironments(entities), cnt, nil
}

// Extensional APIs.
