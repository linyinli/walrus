// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types/object"
)

// ServiceRelationshipQueryInput is the input for the ServiceRelationship query.
type ServiceRelationshipQueryInput struct {
	// ID holds the value of the "id" field.
	ID object.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ServiceRelationshipQueryInput to ServiceRelationship.
func (in ServiceRelationshipQueryInput) Model() *ServiceRelationship {
	return &ServiceRelationship{
		ID: in.ID,
	}
}

// ServiceRelationshipCreateInput is the input for the ServiceRelationship creation.
type ServiceRelationshipCreateInput struct {
	// Service to which the dependency belongs.
	Dependency ServiceQueryInput `json:"dependency"`
}

// Model converts the ServiceRelationshipCreateInput to ServiceRelationship.
func (in ServiceRelationshipCreateInput) Model() *ServiceRelationship {
	var entity = &ServiceRelationship{}
	entity.DependencyID = in.Dependency.ID
	return entity
}

// ServiceRelationshipUpdateInput is the input for the ServiceRelationship modification.
type ServiceRelationshipUpdateInput struct {
	// ID holds the value of the "id" field.
	ID object.ID `uri:"id" json:"-"`
	// Service to which the dependency belongs.
	Dependency ServiceQueryInput `json:"dependency,omitempty"`
}

// Model converts the ServiceRelationshipUpdateInput to ServiceRelationship.
func (in ServiceRelationshipUpdateInput) Model() *ServiceRelationship {
	var entity = &ServiceRelationship{
		ID: in.ID,
	}
	entity.DependencyID = in.Dependency.ID
	return entity
}

// ServiceRelationshipOutput is the output for the ServiceRelationship.
type ServiceRelationshipOutput struct {
	// ID holds the value of the "id" field.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "createTime" field.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// ID list of the service includes all dependencies and the service itself.
	Path []object.ID `json:"path,omitempty"`
	// Type of the relationship.
	Type string `json:"type,omitempty"`
	// Service to which the dependency belongs.
	Dependency *ServiceOutput `json:"dependency,omitempty"`
}

// ExposeServiceRelationship converts the ServiceRelationship to ServiceRelationshipOutput.
func ExposeServiceRelationship(in *ServiceRelationship) *ServiceRelationshipOutput {
	if in == nil {
		return nil
	}
	var entity = &ServiceRelationshipOutput{
		ID:         in.ID,
		CreateTime: in.CreateTime,
		Path:       in.Path,
		Type:       in.Type,
		Dependency: ExposeService(in.Edges.Dependency),
	}
	if in.DependencyID != "" {
		if entity.Dependency == nil {
			entity.Dependency = &ServiceOutput{}
		}
		entity.Dependency.ID = in.DependencyID
	}
	return entity
}

// ExposeServiceRelationships converts the ServiceRelationship slice to ServiceRelationshipOutput pointer slice.
func ExposeServiceRelationships(in []*ServiceRelationship) []*ServiceRelationshipOutput {
	var out = make([]*ServiceRelationshipOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeServiceRelationship(in[i])
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
