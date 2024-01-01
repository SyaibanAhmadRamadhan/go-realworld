// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/ginfra/gdb/generator.go. 

package model

import (
	"errors"
)

// TagTableName this table or collection name
const TagTableName string = "tag"

// NewTag is a struct with pointer that represents the table Tag in the database.
func NewTag() *Tag {
	return &Tag{}
}

// NewTagWithOutPtr is a struct without pointer that represents the table Tag in the database.
func NewTagWithOutPtr() Tag {
	return Tag{}
}

// FieldId is a field or column in the table Tag.
func (t *Tag) FieldId() string {
	return "_id"
}

// SetId is a setter for the field or column Id in the table Tag.
func (t *Tag) SetId(param string) string {
	t.Id = param
	return "_id"
}

// FieldName is a field or column in the table Tag.
func (t *Tag) FieldName() string {
	return "name"
}

// SetName is a setter for the field or column Name in the table Tag.
func (t *Tag) SetName(param string) string {
	t.Name = param
	return "name"
}

// AllField is a function to get all field or column in the table Tag.
func (t *Tag) AllField() (str []string) {
	str = []string{ 
		`_id`,
		`name`,
	}
	return
}

// OrderFields is a function to get all field or column in the table Tag.
func (t *Tag) OrderFields() (str []string) {
	str = []string{ 
		`_id`,
		`name`,
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Tag.
func (t *Tag) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case t.FieldId():
			values = append(values, t.Id)
		case t.FieldName():
			values = append(values, t.Name)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Tag.
func (t *Tag) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case t.FieldId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Id")
			}
			t.SetId(val)
		case t.FieldName():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Name")
			}
			t.SetName(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

