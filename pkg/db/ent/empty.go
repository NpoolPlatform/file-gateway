// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/empty"
)

// Empty is the model entity for the Empty schema.
type Empty struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Empty) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case empty.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Empty", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Empty fields.
func (e *Empty) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case empty.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Empty.
// Note that you need to call Empty.Unwrap() before calling this method if this Empty
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Empty) Update() *EmptyUpdateOne {
	return (&EmptyClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Empty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Empty) Unwrap() *Empty {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Empty is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Empty) String() string {
	var builder strings.Builder
	builder.WriteString("Empty(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Empties is a parsable slice of Empty.
type Empties []*Empty

func (e Empties) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
