// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/pkg/dao/types/property"
)

// ServiceRevisionCreateInput holds the creation input of the ServiceRevision entity.
type ServiceRevisionCreateInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Project     *ProjectQueryInput     `uri:",inline" query:"-" json:"project"`
	Environment *EnvironmentQueryInput `uri:",inline" query:"-" json:"environment"`
	Service     *ServiceQueryInput     `uri:",inline" query:"-" json:"service"`

	Output                    string                      `uri:"-" query:"-" json:"output"`
	InputPlan                 string                      `uri:"-" query:"-" json:"inputPlan"`
	TemplateVersion           string                      `uri:"-" query:"-" json:"templateVersion"`
	TemplateID                string                      `uri:"-" query:"-" json:"templateID"`
	Attributes                property.Values             `uri:"-" query:"-" json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `uri:"-" query:"-" json:"variables,omitempty"`
	DeployerType              string                      `uri:"-" query:"-" json:"deployerType,omitempty"`
	Duration                  int                         `uri:"-" query:"-" json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `uri:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	Tags                      []string                    `uri:"-" query:"-" json:"tags,omitempty"`
}

// Model returns the ServiceRevision entity for creating,
// after validating.
func (srci *ServiceRevisionCreateInput) Model() *ServiceRevision {
	if srci == nil {
		return nil
	}

	sr := &ServiceRevision{
		Output:                    srci.Output,
		InputPlan:                 srci.InputPlan,
		TemplateVersion:           srci.TemplateVersion,
		TemplateID:                srci.TemplateID,
		Attributes:                srci.Attributes,
		Variables:                 srci.Variables,
		DeployerType:              srci.DeployerType,
		Duration:                  srci.Duration,
		PreviousRequiredProviders: srci.PreviousRequiredProviders,
		Tags:                      srci.Tags,
	}

	if srci.Project != nil {
		sr.ProjectID = srci.Project.ID
	}
	if srci.Environment != nil {
		sr.EnvironmentID = srci.Environment.ID
	}
	if srci.Service != nil {
		sr.ServiceID = srci.Service.ID
	}

	return sr
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (srci *ServiceRevisionCreateInput) Load() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.LoadWith(srci.inputConfig.Context, srci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (srci *ServiceRevisionCreateInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if srci.Project != nil {
		err = srci.Project.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}
	if srci.Environment != nil {
		err = srci.Environment.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}
	if srci.Service != nil {
		err = srci.Service.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}

	return nil
}

// ServiceRevisionCreateInputs holds the creation input item of the ServiceRevision entities.
type ServiceRevisionCreateInputsItem struct {
	Output                    string                      `uri:"-" query:"-" json:"output"`
	InputPlan                 string                      `uri:"-" query:"-" json:"inputPlan"`
	TemplateVersion           string                      `uri:"-" query:"-" json:"templateVersion"`
	TemplateID                string                      `uri:"-" query:"-" json:"templateID"`
	Attributes                property.Values             `uri:"-" query:"-" json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `uri:"-" query:"-" json:"variables,omitempty"`
	DeployerType              string                      `uri:"-" query:"-" json:"deployerType,omitempty"`
	Duration                  int                         `uri:"-" query:"-" json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `uri:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	Tags                      []string                    `uri:"-" query:"-" json:"tags,omitempty"`
}

// ServiceRevisionCreateInputs holds the creation input of the ServiceRevision entities.
type ServiceRevisionCreateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Project     *ProjectQueryInput     `uri:",inline" query:"-" json:"project"`
	Environment *EnvironmentQueryInput `uri:",inline" query:"-" json:"environment"`
	Service     *ServiceQueryInput     `uri:",inline" query:"-" json:"service"`

	Items []*ServiceRevisionCreateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the ServiceRevision entities for creating,
// after validating.
func (srci *ServiceRevisionCreateInputs) Model() []*ServiceRevision {
	if srci == nil || len(srci.Items) == 0 {
		return nil
	}

	srs := make([]*ServiceRevision, len(srci.Items))

	for i := range srci.Items {
		sr := &ServiceRevision{
			Output:                    srci.Items[i].Output,
			InputPlan:                 srci.Items[i].InputPlan,
			TemplateVersion:           srci.Items[i].TemplateVersion,
			TemplateID:                srci.Items[i].TemplateID,
			Attributes:                srci.Items[i].Attributes,
			Variables:                 srci.Items[i].Variables,
			DeployerType:              srci.Items[i].DeployerType,
			Duration:                  srci.Items[i].Duration,
			PreviousRequiredProviders: srci.Items[i].PreviousRequiredProviders,
			Tags:                      srci.Items[i].Tags,
		}

		if srci.Project != nil {
			sr.ProjectID = srci.Project.ID
		}
		if srci.Environment != nil {
			sr.EnvironmentID = srci.Environment.ID
		}
		if srci.Service != nil {
			sr.ServiceID = srci.Service.ID
		}

		srs[i] = sr
	}

	return srs
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (srci *ServiceRevisionCreateInputs) Load() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.LoadWith(srci.inputConfig.Context, srci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (srci *ServiceRevisionCreateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if len(srci.Items) == 0 {
		return errors.New("empty items")
	}

	if srci.Project != nil {
		err = srci.Project.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}
	if srci.Environment != nil {
		err = srci.Environment.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}
	if srci.Service != nil {
		err = srci.Service.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}
	return nil
}

// ServiceRevisionDeleteInput holds the deletion input of the ServiceRevision entity.
type ServiceRevisionDeleteInput = ServiceRevisionQueryInput

// ServiceRevisionDeleteInputs holds the deletion input item of the ServiceRevision entities.
type ServiceRevisionDeleteInputsItem struct {
	ID object.ID `uri:"-" query:"-" json:"id"`
}

// ServiceRevisionDeleteInputs holds the deletion input of the ServiceRevision entities.
type ServiceRevisionDeleteInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Project     *ProjectQueryInput     `uri:",inline" query:"-" json:"project"`
	Environment *EnvironmentQueryInput `uri:",inline" query:"-" json:"environment"`
	Service     *ServiceQueryInput     `uri:",inline" query:"-" json:"service"`

	Items []*ServiceRevisionDeleteInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the ServiceRevision entities for deleting,
// after validating.
func (srdi *ServiceRevisionDeleteInputs) Model() []*ServiceRevision {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	srs := make([]*ServiceRevision, len(srdi.Items))
	for i := range srdi.Items {
		srs[i] = &ServiceRevision{
			ID: srdi.Items[i].ID,
		}
	}
	return srs
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (srdi *ServiceRevisionDeleteInputs) Load() error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	return srdi.LoadWith(srdi.inputConfig.Context, srdi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (srdi *ServiceRevisionDeleteInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	if len(srdi.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.ServiceRevisions().Query()

	if srdi.Project != nil {
		err = srdi.Project.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.ProjectID(srdi.Project.ID))
	}

	if srdi.Environment != nil {
		err = srdi.Environment.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.EnvironmentID(srdi.Environment.ID))
	}

	if srdi.Service != nil {
		err = srdi.Service.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.ServiceID(srdi.Service.ID))
	}

	ids := make([]object.ID, 0, len(srdi.Items))

	for i := range srdi.Items {
		if srdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if srdi.Items[i].ID != "" {
			ids = append(ids, srdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(servicerevision.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ServiceRevisionQueryInput holds the query input of the ServiceRevision entity.
type ServiceRevisionQueryInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Project     *ProjectQueryInput     `uri:",inline" query:"-" json:"-"`
	Environment *EnvironmentQueryInput `uri:",inline" query:"-" json:"-"`
	Service     *ServiceQueryInput     `uri:",inline" query:"-" json:"-"`

	Refer *object.Refer `uri:"servicerevision,default=\"\"" query:"-" json:"-"`
	ID    object.ID     `uri:"id" query:"-" json:"id"` // TODO(thxCode): remove the uri:"id" after supporting hierarchical routes.
}

// Model returns the ServiceRevision entity for querying,
// after validating.
func (srqi *ServiceRevisionQueryInput) Model() *ServiceRevision {
	if srqi == nil {
		return nil
	}

	return &ServiceRevision{
		ID: srqi.ID,
	}
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (srqi *ServiceRevisionQueryInput) Load() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.LoadWith(srqi.inputConfig.Context, srqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (srqi *ServiceRevisionQueryInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if srqi.Refer != nil && *srqi.Refer == "" {
		return nil
	}

	q := cs.ServiceRevisions().Query()

	if srqi.Project != nil {
		err = srqi.Project.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.ProjectID(srqi.Project.ID))
	}

	if srqi.Environment != nil {
		err = srqi.Environment.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.EnvironmentID(srqi.Environment.ID))
	}

	if srqi.Service != nil {
		err = srqi.Service.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.ServiceID(srqi.Service.ID))
	}

	if srqi.Refer != nil {
		if srqi.Refer.IsID() {
			q.Where(
				servicerevision.ID(srqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of servicerevision")
		}
	} else if srqi.ID != "" {
		q.Where(
			servicerevision.ID(srqi.ID))
	} else {
		return errors.New("invalid identify of servicerevision")
	}

	srqi.ID, err = q.OnlyID(ctx)
	return err
}

// ServiceRevisionQueryInputs holds the query input of the ServiceRevision entities.
type ServiceRevisionQueryInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Project     *ProjectQueryInput     `uri:",inline" query:"-" json:"project"`
	Environment *EnvironmentQueryInput `uri:",inline" query:"-" json:"environment"`
	Service     *ServiceQueryInput     `uri:",inline" query:"-" json:"service"`
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (srqi *ServiceRevisionQueryInputs) Load() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.LoadWith(srqi.inputConfig.Context, srqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (srqi *ServiceRevisionQueryInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if srqi.Project != nil {
		err = srqi.Project.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}

	if srqi.Environment != nil {
		err = srqi.Environment.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}

	if srqi.Service != nil {
		err = srqi.Service.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
	}

	return err
}

// ServiceRevisionUpdateInput holds the modification input of the ServiceRevision entity.
type ServiceRevisionUpdateInput struct {
	ServiceRevisionQueryInput `uri:",inline" query:"-" json:",inline"`

	TemplateVersion           string                      `uri:"-" query:"-" json:"templateVersion,omitempty"`
	Attributes                property.Values             `uri:"-" query:"-" json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `uri:"-" query:"-" json:"variables,omitempty"`
	InputPlan                 string                      `uri:"-" query:"-" json:"inputPlan,omitempty"`
	Output                    string                      `uri:"-" query:"-" json:"output,omitempty"`
	DeployerType              string                      `uri:"-" query:"-" json:"deployerType,omitempty"`
	Duration                  int                         `uri:"-" query:"-" json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `uri:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	Tags                      []string                    `uri:"-" query:"-" json:"tags,omitempty"`
}

// Model returns the ServiceRevision entity for modifying,
// after validating.
func (srui *ServiceRevisionUpdateInput) Model() *ServiceRevision {
	if srui == nil {
		return nil
	}

	sr := &ServiceRevision{
		ID:                        srui.ID,
		TemplateVersion:           srui.TemplateVersion,
		Attributes:                srui.Attributes,
		Variables:                 srui.Variables,
		InputPlan:                 srui.InputPlan,
		Output:                    srui.Output,
		DeployerType:              srui.DeployerType,
		Duration:                  srui.Duration,
		PreviousRequiredProviders: srui.PreviousRequiredProviders,
		Tags:                      srui.Tags,
	}

	return sr
}

// ServiceRevisionUpdateInputs holds the modification input item of the ServiceRevision entities.
type ServiceRevisionUpdateInputsItem struct {
	ID object.ID `uri:"-" query:"-" json:"id"`

	TemplateVersion           string                      `uri:"-" query:"-" json:"templateVersion,omitempty"`
	Attributes                property.Values             `uri:"-" query:"-" json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `uri:"-" query:"-" json:"variables,omitempty"`
	InputPlan                 string                      `uri:"-" query:"-" json:"inputPlan,omitempty"`
	Output                    string                      `uri:"-" query:"-" json:"output,omitempty"`
	DeployerType              string                      `uri:"-" query:"-" json:"deployerType,omitempty"`
	Duration                  int                         `uri:"-" query:"-" json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `uri:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	Tags                      []string                    `uri:"-" query:"-" json:"tags,omitempty"`
}

// ServiceRevisionUpdateInputs holds the modification input of the ServiceRevision entities.
type ServiceRevisionUpdateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Project     *ProjectQueryInput     `uri:",inline" query:"-" json:"project"`
	Environment *EnvironmentQueryInput `uri:",inline" query:"-" json:"environment"`
	Service     *ServiceQueryInput     `uri:",inline" query:"-" json:"service"`

	Items []*ServiceRevisionUpdateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the ServiceRevision entities for modifying,
// after validating.
func (srui *ServiceRevisionUpdateInputs) Model() []*ServiceRevision {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	srs := make([]*ServiceRevision, len(srui.Items))

	for i := range srui.Items {
		sr := &ServiceRevision{
			ID:                        srui.Items[i].ID,
			TemplateVersion:           srui.Items[i].TemplateVersion,
			Attributes:                srui.Items[i].Attributes,
			Variables:                 srui.Items[i].Variables,
			InputPlan:                 srui.Items[i].InputPlan,
			Output:                    srui.Items[i].Output,
			DeployerType:              srui.Items[i].DeployerType,
			Duration:                  srui.Items[i].Duration,
			PreviousRequiredProviders: srui.Items[i].PreviousRequiredProviders,
			Tags:                      srui.Items[i].Tags,
		}

		srs[i] = sr
	}

	return srs
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (srui *ServiceRevisionUpdateInputs) Load() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.LoadWith(srui.inputConfig.Context, srui.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (srui *ServiceRevisionUpdateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if len(srui.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.ServiceRevisions().Query()

	if srui.Project != nil {
		err = srui.Project.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.ProjectID(srui.Project.ID))
	}

	if srui.Environment != nil {
		err = srui.Environment.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.EnvironmentID(srui.Environment.ID))
	}

	if srui.Service != nil {
		err = srui.Service.LoadWith(ctx, cs)
		if err != nil {
			return err
		}
		q.Where(
			servicerevision.ServiceID(srui.Service.ID))
	}

	ids := make([]object.ID, 0, len(srui.Items))

	for i := range srui.Items {
		if srui.Items[i] == nil {
			return errors.New("nil item")
		}

		if srui.Items[i].ID != "" {
			ids = append(ids, srui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(servicerevision.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ServiceRevisionOutput holds the output of the ServiceRevision entity.
type ServiceRevisionOutput struct {
	ID                        object.ID                   `json:"id,omitempty"`
	CreateTime                *time.Time                  `json:"createTime,omitempty"`
	Status                    string                      `json:"status,omitempty"`
	StatusMessage             string                      `json:"statusMessage,omitempty"`
	TemplateID                string                      `json:"templateID,omitempty"`
	TemplateVersion           string                      `json:"templateVersion,omitempty"`
	Attributes                property.Values             `json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `json:"variables,omitempty"`
	DeployerType              string                      `json:"deployerType,omitempty"`
	Duration                  int                         `json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `json:"previousRequiredProviders,omitempty"`
	Tags                      []string                    `json:"tags,omitempty"`

	Project     *ProjectOutput     `json:"project,omitempty"`
	Environment *EnvironmentOutput `json:"environment,omitempty"`
	Service     *ServiceOutput     `json:"service,omitempty"`
}

// View returns the output of ServiceRevision.
func (sr *ServiceRevision) View() *ServiceRevisionOutput {
	return ExposeServiceRevision(sr)
}

// View returns the output of ServiceRevisions.
func (srs ServiceRevisions) View() []*ServiceRevisionOutput {
	return ExposeServiceRevisions(srs)
}

// ExposeServiceRevision converts the ServiceRevision to ServiceRevisionOutput.
func ExposeServiceRevision(sr *ServiceRevision) *ServiceRevisionOutput {
	if sr == nil {
		return nil
	}

	sro := &ServiceRevisionOutput{
		ID:                        sr.ID,
		CreateTime:                sr.CreateTime,
		Status:                    sr.Status,
		StatusMessage:             sr.StatusMessage,
		TemplateID:                sr.TemplateID,
		TemplateVersion:           sr.TemplateVersion,
		Attributes:                sr.Attributes,
		Variables:                 sr.Variables,
		DeployerType:              sr.DeployerType,
		Duration:                  sr.Duration,
		PreviousRequiredProviders: sr.PreviousRequiredProviders,
		Tags:                      sr.Tags,
	}

	if sr.Edges.Project != nil {
		sro.Project = ExposeProject(sr.Edges.Project)
	} else if sr.ProjectID != "" {
		sro.Project = &ProjectOutput{
			ID: sr.ProjectID,
		}
	}
	if sr.Edges.Environment != nil {
		sro.Environment = ExposeEnvironment(sr.Edges.Environment)
	} else if sr.EnvironmentID != "" {
		sro.Environment = &EnvironmentOutput{
			ID: sr.EnvironmentID,
		}
	}
	if sr.Edges.Service != nil {
		sro.Service = ExposeService(sr.Edges.Service)
	} else if sr.ServiceID != "" {
		sro.Service = &ServiceOutput{
			ID: sr.ServiceID,
		}
	}
	return sro
}

// ExposeServiceRevisions converts the ServiceRevision slice to ServiceRevisionOutput pointer slice.
func ExposeServiceRevisions(srs []*ServiceRevision) []*ServiceRevisionOutput {
	if len(srs) == 0 {
		return nil
	}

	sros := make([]*ServiceRevisionOutput, len(srs))
	for i := range srs {
		sros[i] = ExposeServiceRevision(srs[i])
	}
	return sros
}