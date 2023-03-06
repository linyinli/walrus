// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/perspective"
	"github.com/seal-io/seal/pkg/dao/types"
)

// Perspective is the model entity for the Perspective schema.
type Perspective struct {
	config `json:"-"`
	// ID of the ent.
	ID types.ID `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Name for current perspective
	Name string `json:"name,omitempty"`
	// Start time for current perspective
	StartTime string `json:"startTime,omitempty"`
	// End time for current perspective
	EndTime string `json:"endTime,omitempty"`
	// Is builtin Perspective
	Builtin bool `json:"builtin,omitempty"`
	// Indicated the perspective included allocation queries, record the used query condition
	AllocationQueries []types.QueryCondition `json:"allocationQueries,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Perspective) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case perspective.FieldAllocationQueries:
			values[i] = new([]byte)
		case perspective.FieldBuiltin:
			values[i] = new(sql.NullBool)
		case perspective.FieldName, perspective.FieldStartTime, perspective.FieldEndTime:
			values[i] = new(sql.NullString)
		case perspective.FieldCreateTime, perspective.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case perspective.FieldID:
			values[i] = new(types.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Perspective", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Perspective fields.
func (pe *Perspective) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case perspective.FieldID:
			if value, ok := values[i].(*types.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case perspective.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createTime", values[i])
			} else if value.Valid {
				pe.CreateTime = new(time.Time)
				*pe.CreateTime = value.Time
			}
		case perspective.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updateTime", values[i])
			} else if value.Valid {
				pe.UpdateTime = new(time.Time)
				*pe.UpdateTime = value.Time
			}
		case perspective.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case perspective.FieldStartTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field startTime", values[i])
			} else if value.Valid {
				pe.StartTime = value.String
			}
		case perspective.FieldEndTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field endTime", values[i])
			} else if value.Valid {
				pe.EndTime = value.String
			}
		case perspective.FieldBuiltin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field builtin", values[i])
			} else if value.Valid {
				pe.Builtin = value.Bool
			}
		case perspective.FieldAllocationQueries:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field allocationQueries", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.AllocationQueries); err != nil {
					return fmt.Errorf("unmarshal field allocationQueries: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Perspective.
// Note that you need to call Perspective.Unwrap() before calling this method if this Perspective
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Perspective) Update() *PerspectiveUpdateOne {
	return NewPerspectiveClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Perspective entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Perspective) Unwrap() *Perspective {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("model: Perspective is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Perspective) String() string {
	var builder strings.Builder
	builder.WriteString("Perspective(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	if v := pe.CreateTime; v != nil {
		builder.WriteString("createTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pe.UpdateTime; v != nil {
		builder.WriteString("updateTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("startTime=")
	builder.WriteString(pe.StartTime)
	builder.WriteString(", ")
	builder.WriteString("endTime=")
	builder.WriteString(pe.EndTime)
	builder.WriteString(", ")
	builder.WriteString("builtin=")
	builder.WriteString(fmt.Sprintf("%v", pe.Builtin))
	builder.WriteString(", ")
	builder.WriteString("allocationQueries=")
	builder.WriteString(fmt.Sprintf("%v", pe.AllocationQueries))
	builder.WriteByte(')')
	return builder.String()
}

// Perspectives is a parsable slice of Perspective.
type Perspectives []*Perspective