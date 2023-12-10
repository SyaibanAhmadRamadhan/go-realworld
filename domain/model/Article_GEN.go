package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"
)

// NewArticle is a struct with pointer that represents the table Article in the database.
func NewArticle() *Article {
	return &Article{}
}

// NewArticleWithOutPtr is a struct without pointer that represents the table Article in the database.
func NewArticleWithOutPtr() Article {
	return Article{}
}

// TableName is a function to get table name
func (a *Article) TableName() (table string) {
	return "article"
}

// FieldID is a field or column in the table Article.
func (a *Article) FieldID() string {
	return "_id"
}

// SetID is a setter for the field or column ID in the table Article.
func (a *Article) SetID(param string) {
	a.ID = param
}

// FieldSlug is a field or column in the table Article.
func (a *Article) FieldSlug() string {
	return "slug"
}

// SetSlug is a setter for the field or column Slug in the table Article.
func (a *Article) SetSlug(param string) {
	a.Slug = param
}

// FieldTitle is a field or column in the table Article.
func (a *Article) FieldTitle() string {
	return "title"
}

// SetTitle is a setter for the field or column Title in the table Article.
func (a *Article) SetTitle(param string) {
	a.Title = param
}

// FieldDescription is a field or column in the table Article.
func (a *Article) FieldDescription() string {
	return "description"
}

// SetDescription is a setter for the field or column Description in the table Article.
func (a *Article) SetDescription(param string) {
	a.Description = param
}

// FieldBody is a field or column in the table Article.
func (a *Article) FieldBody() string {
	return "body"
}

// SetBody is a setter for the field or column Body in the table Article.
func (a *Article) SetBody(param string) {
	a.Body = param
}

// FieldCreatedAt is a field or column in the table Article.
func (a *Article) FieldCreatedAt() string {
	return "createdAt"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table Article.
func (a *Article) SetCreatedAt(param string) {
	a.CreatedAt = param
}

// FieldUpdatedAt is a field or column in the table Article.
func (a *Article) FieldUpdatedAt() string {
	return "updatedAt"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table Article.
func (a *Article) SetUpdatedAt(param string) {
	a.UpdatedAt = param
}

// FieldTagID is a field or column in the table Article.
func (a *Article) FieldTagID() string {
	return "tagID"
}

// SetTagID is a setter for the field or column TagID in the table Article.
func (a *Article) SetTagID(param []string) {
	a.TagID = param
}

// FieldAuthorID is a field or column in the table Article.
func (a *Article) FieldAuthorID() string {
	return "authorID"
}

// SetAuthorID is a setter for the field or column AuthorID in the table Article.
func (a *Article) SetAuthorID(param int) {
	a.AuthorID = param
}

// AllField is a function to get all field or column in the table Article.
func (a *Article) AllField() (str []string) {
	str = []string{ 
		`tagID`,
		`authorID`,
		`_id`,
		`slug`,
		`title`,
		`description`,
		`body`,
		`createdAt`,
		`updatedAt`,
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Article.
func (a *Article) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case a.FieldID():
			values = append(values, a.ID)
		case a.FieldSlug():
			values = append(values, a.Slug)
		case a.FieldTitle():
			values = append(values, a.Title)
		case a.FieldDescription():
			values = append(values, a.Description)
		case a.FieldBody():
			values = append(values, a.Body)
		case a.FieldCreatedAt():
			values = append(values, a.CreatedAt)
		case a.FieldUpdatedAt():
			values = append(values, a.UpdatedAt)
		case a.FieldTagID():
			values = append(values, a.TagID)
		case a.FieldAuthorID():
			values = append(values, a.AuthorID)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Article.
func (a *Article) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case a.FieldAuthorID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type int. field AuthorID")
			}
			a.SetAuthorID(val)
		case a.FieldTagID():
			val, ok := value.([]string)
			if !ok {
				return errors.New("invalid type []string. field TagID")
			}
			a.SetTagID(val)
		case a.FieldSlug():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Slug")
			}
			a.SetSlug(val)
		case a.FieldTitle():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Title")
			}
			a.SetTitle(val)
		case a.FieldDescription():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Description")
			}
			a.SetDescription(val)
		case a.FieldBody():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Body")
			}
			a.SetBody(val)
		case a.FieldCreatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedAt")
			}
			a.SetCreatedAt(val)
		case a.FieldUpdatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field UpdatedAt")
			}
			a.SetUpdatedAt(val)
		case a.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			a.SetID(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

