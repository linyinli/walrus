// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/template"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// TemplateCreateInput holds the creation input of the Template entity.
type TemplateCreateInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Source      string            `uri:"-" query:"-" json:"source"`
	Description string            `uri:"-" query:"-" json:"description,omitempty"`
	Icon        string            `uri:"-" query:"-" json:"icon,omitempty"`
	Labels      map[string]string `uri:"-" query:"-" json:"labels,omitempty"`
}

// Model returns the Template entity for creating,
// after validating.
func (tci *TemplateCreateInput) Model() *Template {
	if tci == nil {
		return nil
	}

	t := &Template{
		Source:      tci.Source,
		Description: tci.Description,
		Icon:        tci.Icon,
		Labels:      tci.Labels,
	}

	return t
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (tci *TemplateCreateInput) Load() error {
	if tci == nil {
		return errors.New("nil receiver")
	}

	return tci.LoadWith(tci.inputConfig.Context, tci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (tci *TemplateCreateInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if tci == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// TemplateCreateInputs holds the creation input item of the Template entities.
type TemplateCreateInputsItem struct {
	Source      string            `uri:"-" query:"-" json:"source"`
	Description string            `uri:"-" query:"-" json:"description,omitempty"`
	Icon        string            `uri:"-" query:"-" json:"icon,omitempty"`
	Labels      map[string]string `uri:"-" query:"-" json:"labels,omitempty"`
}

// TemplateCreateInputs holds the creation input of the Template entities.
type TemplateCreateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*TemplateCreateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the Template entities for creating,
// after validating.
func (tci *TemplateCreateInputs) Model() []*Template {
	if tci == nil || len(tci.Items) == 0 {
		return nil
	}

	ts := make([]*Template, len(tci.Items))

	for i := range tci.Items {
		t := &Template{
			Source:      tci.Items[i].Source,
			Description: tci.Items[i].Description,
			Icon:        tci.Items[i].Icon,
			Labels:      tci.Items[i].Labels,
		}

		ts[i] = t
	}

	return ts
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (tci *TemplateCreateInputs) Load() error {
	if tci == nil {
		return errors.New("nil receiver")
	}

	return tci.LoadWith(tci.inputConfig.Context, tci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (tci *TemplateCreateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if tci == nil {
		return errors.New("nil receiver")
	}

	if len(tci.Items) == 0 {
		return errors.New("empty items")
	}

	return nil
}

// TemplateDeleteInput holds the deletion input of the Template entity.
type TemplateDeleteInput = TemplateQueryInput

// TemplateDeleteInputs holds the deletion input item of the Template entities.
type TemplateDeleteInputsItem struct {
	ID string `uri:"-" query:"-" json:"id"`
}

// TemplateDeleteInputs holds the deletion input of the Template entities.
type TemplateDeleteInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*TemplateDeleteInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the Template entities for deleting,
// after validating.
func (tdi *TemplateDeleteInputs) Model() []*Template {
	if tdi == nil || len(tdi.Items) == 0 {
		return nil
	}

	ts := make([]*Template, len(tdi.Items))
	for i := range tdi.Items {
		ts[i] = &Template{
			ID: tdi.Items[i].ID,
		}
	}
	return ts
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (tdi *TemplateDeleteInputs) Load() error {
	if tdi == nil {
		return errors.New("nil receiver")
	}

	return tdi.LoadWith(tdi.inputConfig.Context, tdi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (tdi *TemplateDeleteInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if tdi == nil {
		return errors.New("nil receiver")
	}

	if len(tdi.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Templates().Query()

	ids := make([]string, 0, len(tdi.Items))

	for i := range tdi.Items {
		if tdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if tdi.Items[i].ID != "" {
			ids = append(ids, tdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(template.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// TemplateQueryInput holds the query input of the Template entity.
type TemplateQueryInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Refer *object.Refer `uri:"template,default=\"\"" query:"-" json:"-"`
	ID    string        `uri:"id" query:"-" json:"id"` // TODO(thxCode): remove the uri:"id" after supporting hierarchical routes.
}

// Model returns the Template entity for querying,
// after validating.
func (tqi *TemplateQueryInput) Model() *Template {
	if tqi == nil {
		return nil
	}

	return &Template{
		ID: tqi.ID,
	}
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (tqi *TemplateQueryInput) Load() error {
	if tqi == nil {
		return errors.New("nil receiver")
	}

	return tqi.LoadWith(tqi.inputConfig.Context, tqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (tqi *TemplateQueryInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if tqi == nil {
		return errors.New("nil receiver")
	}

	if tqi.Refer != nil && *tqi.Refer == "" {
		return nil
	}

	q := cs.Templates().Query()

	if tqi.Refer != nil {
		if tqi.Refer.IsString() {
			q.Where(
				template.ID(tqi.Refer.String()))
		} else {
			return errors.New("invalid identify refer of template")
		}
	} else if tqi.ID != "" {
		q.Where(
			template.ID(tqi.ID))
	} else {
		return errors.New("invalid identify of template")
	}

	tqi.ID, err = q.OnlyID(ctx)
	return err
}

// TemplateQueryInputs holds the query input of the Template entities.
type TemplateQueryInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (tqi *TemplateQueryInputs) Load() error {
	if tqi == nil {
		return errors.New("nil receiver")
	}

	return tqi.LoadWith(tqi.inputConfig.Context, tqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (tqi *TemplateQueryInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if tqi == nil {
		return errors.New("nil receiver")
	}

	return err
}

// TemplateUpdateInput holds the modification input of the Template entity.
type TemplateUpdateInput struct {
	TemplateQueryInput `uri:",inline" query:"-" json:",inline"`

	Description string            `uri:"-" query:"-" json:"description,omitempty"`
	Icon        string            `uri:"-" query:"-" json:"icon,omitempty"`
	Labels      map[string]string `uri:"-" query:"-" json:"labels,omitempty"`
	Source      string            `uri:"-" query:"-" json:"source,omitempty"`
}

// Model returns the Template entity for modifying,
// after validating.
func (tui *TemplateUpdateInput) Model() *Template {
	if tui == nil {
		return nil
	}

	t := &Template{
		ID:          tui.ID,
		Description: tui.Description,
		Icon:        tui.Icon,
		Labels:      tui.Labels,
		Source:      tui.Source,
	}

	return t
}

// TemplateUpdateInputs holds the modification input item of the Template entities.
type TemplateUpdateInputsItem struct {
	ID string `uri:"-" query:"-" json:"id"`

	Description string            `uri:"-" query:"-" json:"description,omitempty"`
	Icon        string            `uri:"-" query:"-" json:"icon,omitempty"`
	Labels      map[string]string `uri:"-" query:"-" json:"labels,omitempty"`
	Source      string            `uri:"-" query:"-" json:"source,omitempty"`
}

// TemplateUpdateInputs holds the modification input of the Template entities.
type TemplateUpdateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*TemplateUpdateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the Template entities for modifying,
// after validating.
func (tui *TemplateUpdateInputs) Model() []*Template {
	if tui == nil || len(tui.Items) == 0 {
		return nil
	}

	ts := make([]*Template, len(tui.Items))

	for i := range tui.Items {
		t := &Template{
			ID:          tui.Items[i].ID,
			Description: tui.Items[i].Description,
			Icon:        tui.Items[i].Icon,
			Labels:      tui.Items[i].Labels,
			Source:      tui.Items[i].Source,
		}

		ts[i] = t
	}

	return ts
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (tui *TemplateUpdateInputs) Load() error {
	if tui == nil {
		return errors.New("nil receiver")
	}

	return tui.LoadWith(tui.inputConfig.Context, tui.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (tui *TemplateUpdateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if tui == nil {
		return errors.New("nil receiver")
	}

	if len(tui.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Templates().Query()

	ids := make([]string, 0, len(tui.Items))

	for i := range tui.Items {
		if tui.Items[i] == nil {
			return errors.New("nil item")
		}

		if tui.Items[i].ID != "" {
			ids = append(ids, tui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(template.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// TemplateOutput holds the output of the Template entity.
type TemplateOutput struct {
	ID            string            `json:"id,omitempty"`
	CreateTime    *time.Time        `json:"createTime,omitempty"`
	UpdateTime    *time.Time        `json:"updateTime,omitempty"`
	Status        string            `json:"status,omitempty"`
	StatusMessage string            `json:"statusMessage,omitempty"`
	Description   string            `json:"description,omitempty"`
	Icon          string            `json:"icon,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Source        string            `json:"source,omitempty"`
}

// View returns the output of Template.
func (t *Template) View() *TemplateOutput {
	return ExposeTemplate(t)
}

// View returns the output of Templates.
func (ts Templates) View() []*TemplateOutput {
	return ExposeTemplates(ts)
}

// ExposeTemplate converts the Template to TemplateOutput.
func ExposeTemplate(t *Template) *TemplateOutput {
	if t == nil {
		return nil
	}

	to := &TemplateOutput{
		ID:            t.ID,
		CreateTime:    t.CreateTime,
		UpdateTime:    t.UpdateTime,
		Status:        t.Status,
		StatusMessage: t.StatusMessage,
		Description:   t.Description,
		Icon:          t.Icon,
		Labels:        t.Labels,
		Source:        t.Source,
	}

	return to
}

// ExposeTemplates converts the Template slice to TemplateOutput pointer slice.
func ExposeTemplates(ts []*Template) []*TemplateOutput {
	if len(ts) == 0 {
		return nil
	}

	tos := make([]*TemplateOutput, len(ts))
	for i := range ts {
		tos[i] = ExposeTemplate(ts[i])
	}
	return tos
}