package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"

	"time"
)

// ArticleTableName this table or collection name
const ArticleTableName string = "article"

// NewArticle is a struct with pointer that represents the table Article in the database.
func NewArticle() *Article {
	return &Article{}
}

// NewArticleWithOutPtr is a struct without pointer that represents the table Article in the database.
func NewArticleWithOutPtr() Article {
	return Article{}
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
func (a *Article) SetCreatedAt(param time.Time) {
	a.CreatedAt = param
}

// FieldUpdatedAt is a field or column in the table Article.
func (a *Article) FieldUpdatedAt() string {
	return "updatedAt"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table Article.
func (a *Article) SetUpdatedAt(param time.Time) {
	a.UpdatedAt = param
}

// FieldID is a field or column in the table Article.
func (a *Article) FieldID() string {
	return "_id"
}

// SetID is a setter for the field or column ID in the table Article.
func (a *Article) SetID(param string) {
	a.ID = param
}

// FieldAuthorID is a field or column in the table Article.
func (a *Article) FieldAuthorID() string {
	return "authorID"
}

// SetAuthorID is a setter for the field or column AuthorID in the table Article.
func (a *Article) SetAuthorID(param string) {
	a.AuthorID = param
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

// AllField is a function to get all field or column in the table Article.
func (a *Article) AllField() (str []string) {
	str = []string{ 
		`createdAt`,
		`updatedAt`,
		`_id`,
		`authorID`,
		`slug`,
		`title`,
		`description`,
		`body`,
	}
	return
}

// OrderFields is a function to get all field or column in the table Article.
func (a *Article) OrderFields() (str []string) {
	str = []string{ 
		`updatedAt`,
		`_id`,
		`slug`,
		`title`,
		`createdAt`,
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Article.
func (a *Article) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case a.FieldCreatedAt():
			values = append(values, a.CreatedAt)
		case a.FieldUpdatedAt():
			values = append(values, a.UpdatedAt)
		case a.FieldID():
			values = append(values, a.ID)
		case a.FieldAuthorID():
			values = append(values, a.AuthorID)
		case a.FieldSlug():
			values = append(values, a.Slug)
		case a.FieldTitle():
			values = append(values, a.Title)
		case a.FieldDescription():
			values = append(values, a.Description)
		case a.FieldBody():
			values = append(values, a.Body)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Article.
func (a *Article) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
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
			val, ok := value.(time.Time)
			if !ok {
				return errors.New("invalid type time.Time. field CreatedAt")
			}
			a.SetCreatedAt(val)
		case a.FieldUpdatedAt():
			val, ok := value.(time.Time)
			if !ok {
				return errors.New("invalid type time.Time. field UpdatedAt")
			}
			a.SetUpdatedAt(val)
		case a.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			a.SetID(val)
		case a.FieldAuthorID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field AuthorID")
			}
			a.SetAuthorID(val)
		case a.FieldSlug():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Slug")
			}
			a.SetSlug(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

