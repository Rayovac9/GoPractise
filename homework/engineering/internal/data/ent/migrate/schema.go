// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FaresColumns holds the columns for the "fares" table.
	FaresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "org_airport", Type: field.TypeString},
		{Name: "arr_airport", Type: field.TypeString},
		{Name: "passage_type", Type: field.TypeString},
		{Name: "first_travel_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime", "sqlite3": "datetime"}},
		{Name: "last_travel_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime", "sqlite3": "datetime"}},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(6,2)", "sqlite3": "decimal(6,2)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime", "sqlite3": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime", "sqlite3": "datetime"}},
	}
	// FaresTable holds the schema information for the "fares" table.
	FaresTable = &schema.Table{
		Name:        "fares",
		Columns:     FaresColumns,
		PrimaryKey:  []*schema.Column{FaresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FaresTable,
	}
)

func init() {
}
