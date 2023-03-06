package server

import (
	"context"

	"github.com/seal-io/seal/pkg/dao/model/perspective"
	"github.com/seal-io/seal/pkg/dao/types"

	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
)

func (r *Server) initPerspectives(ctx context.Context, opts initOptions) error {
	var builtin = []*model.Perspective{
		perspectiveAll(),
		perspectiveCluster(),
		perspectiveProject(),
	}

	var creates, err = dao.PerspectiveCreates(opts.ModelClient, builtin...)
	if err != nil {
		return err
	}
	for i := range creates {
		err = creates[i].OnConflictColumns(perspective.FieldName).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func perspectiveAll() *model.Perspective {
	return &model.Perspective{
		Name:      "All",
		StartTime: "now-7d",
		EndTime:   "now",
		Builtin:   true,
		AllocationQueries: []types.QueryCondition{
			// daily cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							IncludeAll: true,
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldDay,
				Paging: types.QueryPagination{
					Page:    1,
					PerPage: 10,
				},
			},
			// per project cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							IncludeAll: true,
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldProject,
				Paging: types.QueryPagination{
					Page:    1,
					PerPage: 10,
				},
			},
			// per cluster cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							IncludeAll: true,
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldClusterName,
				Paging: types.QueryPagination{
					Page:    1,
					PerPage: 10,
				},
			},
		},
	}
}

func perspectiveCluster() *model.Perspective {
	return &model.Perspective{
		Name:      "Cluster",
		StartTime: "now-7d",
		EndTime:   "now",
		Builtin:   true,
		AllocationQueries: []types.QueryCondition{
			// daily cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							FieldName: types.FilterFieldConnectorID,
							Operator:  types.OperatorIn,
							Values:    []string{"${connectorID}"},
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldDay,
				Step:    types.StepDay,
				Paging: types.QueryPagination{
					Page:    1,
					PerPage: 10,
				},
			},
			// per namespace cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							IncludeAll: true,
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldNamespace,
				Paging: types.QueryPagination{
					Page:    1,
					PerPage: 10,
				},
			},
			// workload per day cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							IncludeAll: true,
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldWorkload,
				Paging: types.QueryPagination{
					Page:    1,
					PerPage: 10,
				},
				Step: types.StepDay,
			},
		},
	}
}

func perspectiveProject() *model.Perspective {
	return &model.Perspective{
		Name:      "Project",
		StartTime: "now-7d",
		EndTime:   "now",
		Builtin:   true,
		AllocationQueries: []types.QueryCondition{
			// app cost
			{
				Filters: types.AllocationCostFilters{
					{
						{
							FieldName: types.FilterFieldProject,
							Operator:  types.OperatorIn,
							Values:    []string{"${project}"},
						},
					},
				},
				SharedCosts: types.ShareCosts{
					{
						IdleCostFilters: types.IdleCostFilters{
							{
								IncludeAll: true,
							},
						},
						ManagementCostFilters: types.ManagementCostFilters{
							{
								IncludeAll: true,
							},
						},
						SharingStrategy: types.SharingStrategyProportionally,
					},
				},
				GroupBy: types.GroupByFieldApplication,
				Step:    types.StepDay,
			},
		},
	}
}