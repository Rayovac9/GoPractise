// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/webmin7761/go-school/homework/engineering/internal/data/ent/fare"
)

// Fare is the model entity for the Fare schema.
type Fare struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// OrgAirport holds the value of the "org_airport" field.
	OrgAirport string `json:"org_airport,omitempty"`
	// ArrAirport holds the value of the "arr_airport" field.
	ArrAirport string `json:"arr_airport,omitempty"`
	// PassageType holds the value of the "passage_type" field.
	PassageType string `json:"passage_type,omitempty"`
	// FirstTravelDate holds the value of the "first_travel_date" field.
	FirstTravelDate time.Time `json:"first_travel_date,omitempty"`
	// LastTravelDate holds the value of the "last_travel_date" field.
	LastTravelDate time.Time `json:"last_travel_date,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Fare) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fare.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case fare.FieldID:
			values[i] = new(sql.NullInt64)
		case fare.FieldOrgAirport, fare.FieldArrAirport, fare.FieldPassageType:
			values[i] = new(sql.NullString)
		case fare.FieldFirstTravelDate, fare.FieldLastTravelDate, fare.FieldCreatedAt, fare.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Fare", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Fare fields.
func (f *Fare) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fare.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int64(value.Int64)
		case fare.FieldOrgAirport:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_airport", values[i])
			} else if value.Valid {
				f.OrgAirport = value.String
			}
		case fare.FieldArrAirport:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arr_airport", values[i])
			} else if value.Valid {
				f.ArrAirport = value.String
			}
		case fare.FieldPassageType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passage_type", values[i])
			} else if value.Valid {
				f.PassageType = value.String
			}
		case fare.FieldFirstTravelDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field first_travel_date", values[i])
			} else if value.Valid {
				f.FirstTravelDate = value.Time
			}
		case fare.FieldLastTravelDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_travel_date", values[i])
			} else if value.Valid {
				f.LastTravelDate = value.Time
			}
		case fare.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				f.Amount = value.Float64
			}
		case fare.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case fare.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Fare.
// Note that you need to call Fare.Unwrap() before calling this method if this Fare
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Fare) Update() *FareUpdateOne {
	return (&FareClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Fare entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Fare) Unwrap() *Fare {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Fare is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Fare) String() string {
	var builder strings.Builder
	builder.WriteString("Fare(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", org_airport=")
	builder.WriteString(f.OrgAirport)
	builder.WriteString(", arr_airport=")
	builder.WriteString(f.ArrAirport)
	builder.WriteString(", passage_type=")
	builder.WriteString(f.PassageType)
	builder.WriteString(", first_travel_date=")
	builder.WriteString(f.FirstTravelDate.Format(time.ANSIC))
	builder.WriteString(", last_travel_date=")
	builder.WriteString(f.LastTravelDate.Format(time.ANSIC))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", f.Amount))
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Fares is a parsable slice of Fare.
type Fares []*Fare

func (f Fares) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
