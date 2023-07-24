// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types/object"
)

// ProjectQueryInput is the input for the Project query.
type ProjectQueryInput struct {
	// ID holds the value of the "id" field.
	ID object.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ProjectQueryInput to Project.
func (in ProjectQueryInput) Model() *Project {
	return &Project{
		ID: in.ID,
	}
}

// ProjectCreateInput is the input for the Project creation.
type ProjectCreateInput struct {
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
}

// Model converts the ProjectCreateInput to Project.
func (in ProjectCreateInput) Model() *Project {
	var entity = &Project{
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
	}
	return entity
}

// ProjectUpdateInput is the input for the Project modification.
type ProjectUpdateInput struct {
	// ID holds the value of the "id" field.
	ID object.ID `uri:"id" json:"-"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
}

// Model converts the ProjectUpdateInput to Project.
func (in ProjectUpdateInput) Model() *Project {
	var entity = &Project{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
	}
	return entity
}

// ProjectOutput is the output for the Project.
type ProjectOutput struct {
	// ID holds the value of the "id" field.
	ID object.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// CreateTime holds the value of the "createTime" field.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// UpdateTime holds the value of the "updateTime" field.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
}

// ExposeProject converts the Project to ProjectOutput.
func ExposeProject(in *Project) *ProjectOutput {
	if in == nil {
		return nil
	}
	var entity = &ProjectOutput{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
	}
	return entity
}

// ExposeProjects converts the Project slice to ProjectOutput pointer slice.
func ExposeProjects(in []*Project) []*ProjectOutput {
	var out = make([]*ProjectOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeProject(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
