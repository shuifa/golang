package schema

import (
    "go/ast"
    "reflect"

    "github.com/oushuifa/golang/better/orm/dialect"
)

// Field represents a column of database
type Field struct {
    Name string
    Type string
    Tag  string
}

// Schema represents a table of database
type Schema struct {
    Model      interface{}
    Name       string
    Fields     []*Field
    FieldNames []string
    fieldMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
    return s.fieldMap[name]
}


func Parse(dest interface{}, d dialect.Dialect) *Schema {

    modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

    schema := &Schema{
        Model:      modelType,
        Name:       modelType.Name(),
        fieldMap:   make(map[string]* Field),
    }

    for i := 0; i < modelType.NumField(); i++ {
        f := modelType.Field(i)

        if !f.Anonymous && ast.IsExported(f.Name) {
            field := &Field{
                Name: f.Name,
                Type: d.DataTypeOf(reflect.Indirect(reflect.New(f.Type))),
            }
            if tag, ok := f.Tag.Lookup("orm"); ok {
                field.Tag = tag
            }
            schema.Fields = append(schema.Fields, field)
            schema.fieldMap[field.Name] = field
            schema.FieldNames = append(schema.FieldNames, field.Name)
        }
    }

    return schema
}