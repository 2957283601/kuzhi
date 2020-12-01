// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/kallydev/privacy/ent/jdmodel"
)

// JDModel is the model entity for the JDModel schema.
type JDModel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// IDNumber holds the value of the "id_number" field.
	IDNumber string `json:"id_number,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber int64 `json:"phone_number,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JDModel) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // nickname
		&sql.NullString{}, // password
		&sql.NullString{}, // email
		&sql.NullString{}, // id_number
		&sql.NullInt64{},  // phone_number
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JDModel fields.
func (jm *JDModel) assignValues(values ...interface{}) error {
	if m, n := len(values), len(jdmodel.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	jm.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		jm.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nickname", values[1])
	} else if value.Valid {
		jm.Nickname = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		jm.Password = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[3])
	} else if value.Valid {
		jm.Email = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id_number", values[4])
	} else if value.Valid {
		jm.IDNumber = value.String
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field phone_number", values[5])
	} else if value.Valid {
		jm.PhoneNumber = value.Int64
	}
	return nil
}

// Update returns a builder for updating this JDModel.
// Note that, you need to call JDModel.Unwrap() before calling this method, if this JDModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (jm *JDModel) Update() *JDModelUpdateOne {
	return (&JDModelClient{config: jm.config}).UpdateOne(jm)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (jm *JDModel) Unwrap() *JDModel {
	tx, ok := jm.config.driver.(*txDriver)
	if !ok {
		panic("ent: JDModel is not a transactional entity")
	}
	jm.config.driver = tx.drv
	return jm
}

// String implements the fmt.Stringer.
func (jm *JDModel) String() string {
	var builder strings.Builder
	builder.WriteString("JDModel(")
	builder.WriteString(fmt.Sprintf("id=%v", jm.ID))
	builder.WriteString(", name=")
	builder.WriteString(jm.Name)
	builder.WriteString(", nickname=")
	builder.WriteString(jm.Nickname)
	builder.WriteString(", password=")
	builder.WriteString(jm.Password)
	builder.WriteString(", email=")
	builder.WriteString(jm.Email)
	builder.WriteString(", id_number=")
	builder.WriteString(jm.IDNumber)
	builder.WriteString(", phone_number=")
	builder.WriteString(fmt.Sprintf("%v", jm.PhoneNumber))
	builder.WriteByte(')')
	return builder.String()
}

// JDModels is a parsable slice of JDModel.
type JDModels []*JDModel

func (jm JDModels) config(cfg config) {
	for _i := range jm {
		jm[_i].config = cfg
	}
}
