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

// FieldAuthorId is a field or column in the table Article.
func (a *Article) FieldAuthorId() string {
	return "authorId"
}

// SetAuthorId is a setter for the field or column AuthorId in the table Article.
func (a *Article) SetAuthorId(param string) {
	a.AuthorId = param
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

// FieldId is a field or column in the table Article.
func (a *Article) FieldId() string {
	return "_id"
}

// SetId is a setter for the field or column Id in the table Article.
func (a *Article) SetId(param string) {
	a.Id = param
}

// AllField is a function to get all field or column in the table Article.
func (a *Article) AllField() (str []string) {
	str = []string{ 
		`body`,
		`createdAt`,
		`updatedAt`,
		`_id`,
		`authorId`,
		`slug`,
		`title`,
		`description`,
	}
	return
}

// OrderFields is a function to get all field or column in the table Article.
func (a *Article) OrderFields() (str []string) {
	str = []string{ 
		`createdAt`,
		`updatedAt`,
		`_id`,
		`slug`,
		`title`,
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Article.
func (a *Article) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case a.FieldUpdatedAt():
			values = append(values, a.UpdatedAt)
		case a.FieldId():
			values = append(values, a.Id)
		case a.FieldAuthorId():
			values = append(values, a.AuthorId)
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
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Article.
func (a *Article) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
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
		case a.FieldId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Id")
			}
			a.SetId(val)
		case a.FieldAuthorId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field AuthorId")
			}
			a.SetAuthorId(val)
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
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

