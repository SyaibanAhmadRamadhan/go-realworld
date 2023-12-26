package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"

	"time"
)

// CommentTableName this table or collection name
const CommentTableName string = "comment"

// NewComment is a struct with pointer that represents the table Comment in the database.
func NewComment() *Comment {
	return &Comment{}
}

// NewCommentWithOutPtr is a struct without pointer that represents the table Comment in the database.
func NewCommentWithOutPtr() Comment {
	return Comment{}
}

// FieldId is a field or column in the table Comment.
func (c *Comment) FieldId() string {
	return "_id"
}

// SetId is a setter for the field or column Id in the table Comment.
func (c *Comment) SetId(param string) {
	c.Id = param
}

// FieldArticleId is a field or column in the table Comment.
func (c *Comment) FieldArticleId() string {
	return "articleId"
}

// SetArticleId is a setter for the field or column ArticleId in the table Comment.
func (c *Comment) SetArticleId(param string) {
	c.ArticleId = param
}

// FieldAuthorId is a field or column in the table Comment.
func (c *Comment) FieldAuthorId() string {
	return "authorId"
}

// SetAuthorId is a setter for the field or column AuthorId in the table Comment.
func (c *Comment) SetAuthorId(param string) {
	c.AuthorId = param
}

// FieldBody is a field or column in the table Comment.
func (c *Comment) FieldBody() string {
	return "body"
}

// SetBody is a setter for the field or column Body in the table Comment.
func (c *Comment) SetBody(param string) {
	c.Body = param
}

// FieldCreatedAt is a field or column in the table Comment.
func (c *Comment) FieldCreatedAt() string {
	return "createdAt"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table Comment.
func (c *Comment) SetCreatedAt(param time.Time) {
	c.CreatedAt = param
}

// FieldUpdatedAt is a field or column in the table Comment.
func (c *Comment) FieldUpdatedAt() string {
	return "updatedAt"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table Comment.
func (c *Comment) SetUpdatedAt(param time.Time) {
	c.UpdatedAt = param
}

// AllField is a function to get all field or column in the table Comment.
func (c *Comment) AllField() (str []string) {
	str = []string{ 
		`_id`,
		`articleId`,
		`authorId`,
		`body`,
		`createdAt`,
		`updatedAt`,
	}
	return
}

// OrderFields is a function to get all field or column in the table Comment.
func (c *Comment) OrderFields() (str []string) {
	str = []string{ 
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Comment.
func (c *Comment) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case c.FieldId():
			values = append(values, c.Id)
		case c.FieldArticleId():
			values = append(values, c.ArticleId)
		case c.FieldAuthorId():
			values = append(values, c.AuthorId)
		case c.FieldBody():
			values = append(values, c.Body)
		case c.FieldCreatedAt():
			values = append(values, c.CreatedAt)
		case c.FieldUpdatedAt():
			values = append(values, c.UpdatedAt)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Comment.
func (c *Comment) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case c.FieldId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Id")
			}
			c.SetId(val)
		case c.FieldArticleId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ArticleId")
			}
			c.SetArticleId(val)
		case c.FieldAuthorId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field AuthorId")
			}
			c.SetAuthorId(val)
		case c.FieldBody():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Body")
			}
			c.SetBody(val)
		case c.FieldCreatedAt():
			val, ok := value.(time.Time)
			if !ok {
				return errors.New("invalid type time.Time. field CreatedAt")
			}
			c.SetCreatedAt(val)
		case c.FieldUpdatedAt():
			val, ok := value.(time.Time)
			if !ok {
				return errors.New("invalid type time.Time. field UpdatedAt")
			}
			c.SetUpdatedAt(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

