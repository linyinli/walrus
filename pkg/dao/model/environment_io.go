// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types/object"
)

// EnvironmentQueryInput is the input for the Environment query.
type EnvironmentQueryInput struct {
	// ID holds the value of the "id" field.
	ID object.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the EnvironmentQueryInput to Environment.
func (in EnvironmentQueryInput) Model() *Environment {
	return &Environment{
		ID: in.ID,
	}
}

// EnvironmentCreateInput is the input for the Environment creation.
type EnvironmentCreateInput struct {
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Project to which the environment belongs.
	Project ProjectQueryInput `json:"project"`
	// Connectors holds the value of the connectors edge.
	Connectors []*EnvironmentConnectorRelationshipCreateInput `json:"connectors,omitempty"`
}

// Model converts the EnvironmentCreateInput to Environment.
func (in EnvironmentCreateInput) Model() *Environment {
	var entity = &Environment{
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
	}
	entity.ProjectID = in.Project.ID
	for i := 0; i < len(in.Connectors); i++ {
		if in.Connectors[i] == nil {
			continue
		}
		entity.Edges.Connectors = append(entity.Edges.Connectors, in.Connectors[i].Model())
	}
	return entity
}

// EnvironmentUpdateInput is the input for the Environment modification.
type EnvironmentUpdateInput struct {
	// ID holds the value of the "id" field.
	ID object.ID `uri:"id" json:"-"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Connectors holds the value of the connectors edge.
	Connectors []*EnvironmentConnectorRelationshipUpdateInput `json:"connectors,omitempty"`
}

// Model converts the EnvironmentUpdateInput to Environment.
func (in EnvironmentUpdateInput) Model() *Environment {
	var entity = &Environment{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
	}
	for i := 0; i < len(in.Connectors); i++ {
		if in.Connectors[i] == nil {
			continue
		}
		entity.Edges.Connectors = append(entity.Edges.Connectors, in.Connectors[i].Model())
	}
	return entity
}

// EnvironmentOutput is the output for the Environment.
type EnvironmentOutput struct {
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
	// Project to which the environment belongs.
	Project *ProjectOutput `json:"project,omitempty"`
	// Connectors holds the value of the connectors edge.
	Connectors []*EnvironmentConnectorRelationshipOutput `json:"connectors,omitempty"`
}

// ExposeEnvironment converts the Environment to EnvironmentOutput.
func ExposeEnvironment(in *Environment) *EnvironmentOutput {
	if in == nil {
		return nil
	}
	var entity = &EnvironmentOutput{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
		Project:     ExposeProject(in.Edges.Project),
		Connectors:  ExposeEnvironmentConnectorRelationships(in.Edges.Connectors),
	}
	if in.ProjectID != "" {
		if entity.Project == nil {
			entity.Project = &ProjectOutput{}
		}
		entity.Project.ID = in.ProjectID
	}
	return entity
}

// ExposeEnvironments converts the Environment slice to EnvironmentOutput pointer slice.
func ExposeEnvironments(in []*Environment) []*EnvironmentOutput {
	var out = make([]*EnvironmentOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeEnvironment(in[i])
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
