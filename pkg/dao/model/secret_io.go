// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// SecretQueryInput is the input for the Secret query.
type SecretQueryInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the SecretQueryInput to Secret.
func (in SecretQueryInput) Model() *Secret {
	return &Secret{
		ID: in.ID,
	}
}

// SecretCreateInput is the input for the Secret creation.
type SecretCreateInput struct {
	// The name of secret.
	Name string `json:"name"`
	// The value of secret, store in string.
	Value crypto.String `json:"value,omitempty"`
	// Project to which this secret belongs.
	Project *ProjectQueryInput `json:"project,omitempty"`
}

// Model converts the SecretCreateInput to Secret.
func (in SecretCreateInput) Model() *Secret {
	var entity = &Secret{
		Name:  in.Name,
		Value: in.Value,
	}
	if in.Project != nil {
		entity.ProjectID = in.Project.ID
	}
	return entity
}

// SecretUpdateInput is the input for the Secret modification.
type SecretUpdateInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id" json:"-"`
	// The value of secret, store in string.
	Value crypto.String `json:"value,omitempty"`
}

// Model converts the SecretUpdateInput to Secret.
func (in SecretUpdateInput) Model() *Secret {
	var entity = &Secret{
		ID:    in.ID,
		Value: in.Value,
	}
	return entity
}

// SecretOutput is the output for the Secret.
type SecretOutput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// The name of secret.
	Name string `json:"name,omitempty"`
	// Project to which this secret belongs.
	Project *ProjectOutput `json:"project,omitempty"`
}

// ExposeSecret converts the Secret to SecretOutput.
func ExposeSecret(in *Secret) *SecretOutput {
	if in == nil {
		return nil
	}
	var entity = &SecretOutput{
		ID:         in.ID,
		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
		Name:       in.Name,
		Project:    ExposeProject(in.Edges.Project),
	}
	if entity.Project == nil {
		entity.Project = &ProjectOutput{}
	}
	entity.Project.ID = in.ProjectID
	return entity
}

// ExposeSecrets converts the Secret slice to SecretOutput pointer slice.
func ExposeSecrets(in []*Secret) []*SecretOutput {
	var out = make([]*SecretOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeSecret(in[i])
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