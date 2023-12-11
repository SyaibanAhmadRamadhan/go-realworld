package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"
)

// NewTag is a struct with pointer that represents the table Tag in the database.
func NewTag() *Tag {
	return &Tag{}
}

// NewTagWithOutPtr is a struct without pointer that represents the table Tag in the database.
func NewTagWithOutPtr() Tag {
	return Tag{}
}

// TableName is a function to get table name
func (t *Tag) TableName() (table string) {
	return "tag"
}

// FieldID is a field or column in the table Tag.
func (t *Tag) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table Tag.
func (t *Tag) SetID(param string) {
	t.ID = param
}

// FieldName is a field or column in the table Tag.
func (t *Tag) FieldName() string {
	return "name"
}

// SetName is a setter for the field or column Name in the table Tag.
func (t *Tag) SetName(param string) {
	t.Name = param
}

// AllField is a function to get all field or column in the table Tag.
func (t *Tag) AllField() (str []string) {
	str = []string{ 
		`id`,
		`name`,
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Tag.
func (t *Tag) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case t.FieldID():
			values = append(values, t.ID)
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
		case t.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			t.SetID(val)
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

