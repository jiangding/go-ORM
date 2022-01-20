package schema

import (
	"go-ORM/dialect"
	"testing"
)

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age int
	Weight int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)

	if schema.Name != "User" || len(schema.Fields) != 3 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}

func TestSchema_RecordValues(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	values := schema.RecordValues(&User{"Tom", 18, 33})

	name := values[0].(string)
	age := values[1].(int)

	if name != "Tom" || age != 18 {
		t.Fatal("failed to get values")
	}
}

