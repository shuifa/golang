package schema

import (
    "fmt"
    "testing"

    "github.com/oushuifa/golang/better/orm/dialect"
)

type User struct {
	Name     string `orm:"PRIMARY KEY"`
	Password string `orm:"VARCHAR 64"`
	Id       uint32
}
var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {


	schema := Parse(&User{}, TestDial)
    fmt.Printf("%#v \n", schema)
    if schema.Name != "User" || len(schema.Fields) != 3 {
        t.Fatal("failed to parse User struct")
    }

    if schema.GetField("Name").Tag != "PRIMARY KEY" {
        t.Fatal("failed to parse primary key")
    }
}
