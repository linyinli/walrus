// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// TemplateVersion is the model entity for the TemplateVersion schema.
type TemplateVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// ID of the template.
	TemplateID object.ID `json:"template_id,omitempty"`
	// Name of the template.
	Name string `json:"name,omitempty"`
	// Version of the template.
	Version string `json:"version,omitempty"`
	// Source of the template.
	Source string `json:"source,omitempty"`
	// Schema of the template.
	Schema *types.TemplateSchema `json:"schema,omitempty"`
	// ID of the project to belong, empty means for all projects.
	ProjectID object.ID `json:"project_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TemplateVersionQuery when eager-loading is set.
	Edges        TemplateVersionEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// TemplateVersionEdges holds the relations/edges for other nodes in the graph.
type TemplateVersionEdges struct {
	// Template to which the template version belongs.
	Template *Template `json:"template,omitempty"`
	// Services that belong to the template version.
	Services []*Service `json:"services,omitempty"`
	// Project to which the template version belongs.
	Project *Project `json:"project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TemplateVersionEdges) TemplateOrErr() (*Template, error) {
	if e.loadedTypes[0] {
		if e.Template == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: template.Label}
		}
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// ServicesOrErr returns the Services value or an error if the edge
// was not loaded in eager-loading.
func (e TemplateVersionEdges) ServicesOrErr() ([]*Service, error) {
	if e.loadedTypes[1] {
		return e.Services, nil
	}
	return nil, &NotLoadedError{edge: "services"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TemplateVersionEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[2] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TemplateVersion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case templateversion.FieldSchema:
			values[i] = new([]byte)
		case templateversion.FieldID, templateversion.FieldTemplateID, templateversion.FieldProjectID:
			values[i] = new(object.ID)
		case templateversion.FieldName, templateversion.FieldVersion, templateversion.FieldSource:
			values[i] = new(sql.NullString)
		case templateversion.FieldCreateTime, templateversion.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TemplateVersion fields.
func (tv *TemplateVersion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case templateversion.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tv.ID = *value
			}
		case templateversion.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				tv.CreateTime = new(time.Time)
				*tv.CreateTime = value.Time
			}
		case templateversion.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				tv.UpdateTime = new(time.Time)
				*tv.UpdateTime = value.Time
			}
		case templateversion.FieldTemplateID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field template_id", values[i])
			} else if value != nil {
				tv.TemplateID = *value
			}
		case templateversion.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tv.Name = value.String
			}
		case templateversion.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				tv.Version = value.String
			}
		case templateversion.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				tv.Source = value.String
			}
		case templateversion.FieldSchema:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field schema", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tv.Schema); err != nil {
					return fmt.Errorf("unmarshal field schema: %w", err)
				}
			}
		case templateversion.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				tv.ProjectID = *value
			}
		default:
			tv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TemplateVersion.
// This includes values selected through modifiers, order, etc.
func (tv *TemplateVersion) Value(name string) (ent.Value, error) {
	return tv.selectValues.Get(name)
}

// QueryTemplate queries the "template" edge of the TemplateVersion entity.
func (tv *TemplateVersion) QueryTemplate() *TemplateQuery {
	return NewTemplateVersionClient(tv.config).QueryTemplate(tv)
}

// QueryServices queries the "services" edge of the TemplateVersion entity.
func (tv *TemplateVersion) QueryServices() *ServiceQuery {
	return NewTemplateVersionClient(tv.config).QueryServices(tv)
}

// QueryProject queries the "project" edge of the TemplateVersion entity.
func (tv *TemplateVersion) QueryProject() *ProjectQuery {
	return NewTemplateVersionClient(tv.config).QueryProject(tv)
}

// Update returns a builder for updating this TemplateVersion.
// Note that you need to call TemplateVersion.Unwrap() before calling this method if this TemplateVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (tv *TemplateVersion) Update() *TemplateVersionUpdateOne {
	return NewTemplateVersionClient(tv.config).UpdateOne(tv)
}

// Unwrap unwraps the TemplateVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tv *TemplateVersion) Unwrap() *TemplateVersion {
	_tx, ok := tv.config.driver.(*txDriver)
	if !ok {
		panic("model: TemplateVersion is not a transactional entity")
	}
	tv.config.driver = _tx.drv
	return tv
}

// String implements the fmt.Stringer.
func (tv *TemplateVersion) String() string {
	var builder strings.Builder
	builder.WriteString("TemplateVersion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tv.ID))
	if v := tv.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := tv.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("template_id=")
	builder.WriteString(fmt.Sprintf("%v", tv.TemplateID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tv.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(tv.Version)
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(tv.Source)
	builder.WriteString(", ")
	builder.WriteString("schema=")
	builder.WriteString(fmt.Sprintf("%v", tv.Schema))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", tv.ProjectID))
	builder.WriteByte(')')
	return builder.String()
}

// TemplateVersions is a parsable slice of TemplateVersion.
type TemplateVersions []*TemplateVersion
