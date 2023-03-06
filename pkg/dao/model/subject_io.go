// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/schema"
	"github.com/seal-io/seal/pkg/dao/types"
)

// SubjectQueryInput is the input for the Subject query.
type SubjectQueryInput struct {
	// ID holds the value of the "id" field.
	ID types.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the SubjectQueryInput to Subject.
func (in SubjectQueryInput) Model() *Subject {
	return &Subject{
		ID: in.ID,
	}
}

// SubjectCreateInput is the input for the Subject creation.
type SubjectCreateInput struct {
	// The kind of the subject.
	Kind string `json:"kind,omitempty"`
	// The group of the subject.
	Group string `json:"group,omitempty"`
	// The name of the subject.
	Name string `json:"name"`
	// The detail of the subject.
	Description string `json:"description,omitempty"`
	// Indicate whether the user mount to the group.
	MountTo *bool `json:"mountTo,omitempty"`
	// Indicate whether the user login to the group.
	LoginTo *bool `json:"loginTo,omitempty"`
	// The role list of the subject.
	Roles schema.SubjectRoles `json:"roles,omitempty"`
	// The path of the subject from the root group to itself.
	Paths []string `json:"paths,omitempty"`
	// Indicate whether the subject is builtin.
	Builtin bool `json:"builtin,omitempty"`
}

// Model converts the SubjectCreateInput to Subject.
func (in SubjectCreateInput) Model() *Subject {
	var entity = &Subject{
		Kind:        in.Kind,
		Group:       in.Group,
		Name:        in.Name,
		Description: in.Description,
		MountTo:     in.MountTo,
		LoginTo:     in.LoginTo,
		Roles:       in.Roles,
		Paths:       in.Paths,
		Builtin:     in.Builtin,
	}
	return entity
}

// SubjectUpdateInput is the input for the Subject modification.
type SubjectUpdateInput struct {
	// ID holds the value of the "id" field.
	ID types.ID `uri:"id" json:"-"`
	// The group of the subject.
	Group string `json:"group,omitempty"`
	// The detail of the subject.
	Description string `json:"description,omitempty"`
	// Indicate whether the user mount to the group.
	MountTo *bool `json:"mountTo,omitempty"`
	// Indicate whether the user login to the group.
	LoginTo *bool `json:"loginTo,omitempty"`
	// The role list of the subject.
	Roles schema.SubjectRoles `json:"roles,omitempty"`
	// The path of the subject from the root group to itself.
	Paths []string `json:"paths,omitempty"`
	// Indicate whether the subject is builtin.
	Builtin bool `json:"builtin,omitempty"`
}

// Model converts the SubjectUpdateInput to Subject.
func (in SubjectUpdateInput) Model() *Subject {
	var entity = &Subject{
		ID:          in.ID,
		Group:       in.Group,
		Description: in.Description,
		MountTo:     in.MountTo,
		LoginTo:     in.LoginTo,
		Roles:       in.Roles,
		Paths:       in.Paths,
		Builtin:     in.Builtin,
	}
	return entity
}

// SubjectOutput is the output for the Subject.
type SubjectOutput struct {
	// ID holds the value of the "id" field.
	ID types.ID `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// The kind of the subject.
	Kind string `json:"kind,omitempty"`
	// The group of the subject.
	Group string `json:"group,omitempty"`
	// The name of the subject.
	Name string `json:"name,omitempty"`
	// The detail of the subject.
	Description string `json:"description,omitempty"`
	// Indicate whether the user mount to the group.
	MountTo *bool `json:"mountTo,omitempty"`
	// Indicate whether the user login to the group.
	LoginTo *bool `json:"loginTo,omitempty"`
	// The role list of the subject.
	Roles schema.SubjectRoles `json:"roles,omitempty"`
	// The path of the subject from the root group to itself.
	Paths []string `json:"paths,omitempty"`
	// Indicate whether the subject is builtin.
	Builtin bool `json:"builtin,omitempty"`
}

// ExposeSubject converts the Subject to SubjectOutput.
func ExposeSubject(in *Subject) *SubjectOutput {
	if in == nil {
		return nil
	}
	var entity = &SubjectOutput{
		ID:          in.ID,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
		Kind:        in.Kind,
		Group:       in.Group,
		Name:        in.Name,
		Description: in.Description,
		MountTo:     in.MountTo,
		LoginTo:     in.LoginTo,
		Roles:       in.Roles,
		Paths:       in.Paths,
		Builtin:     in.Builtin,
	}
	return entity
}

// ExposeSubjects converts the Subject slice to SubjectOutput pointer slice.
func ExposeSubjects(in []*Subject) []*SubjectOutput {
	var out = make([]*SubjectOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeSubject(in[i])
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