package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"
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

// FieldCreatedAt is a field or column in the table Comment.
func (c *Comment) FieldCreatedAt() string {
	return "createdAt"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table Comment.
func (c *Comment) SetCreatedAt(param string) {
	c.CreatedAt = param
}

// FieldUpdatedAt is a field or column in the table Comment.
func (c *Comment) FieldUpdatedAt() string {
	return "updatedAt"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table Comment.
func (c *Comment) SetUpdatedAt(param string) {
	c.UpdatedAt = param
}

// FieldID is a field or column in the table Comment.
func (c *Comment) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table Comment.
func (c *Comment) SetID(param int) {
	c.ID = param
}

// FieldArticleID is a field or column in the table Comment.
func (c *Comment) FieldArticleID() string {
	return "articleID"
}

// SetArticleID is a setter for the field or column ArticleID in the table Comment.
func (c *Comment) SetArticleID(param int) {
	c.ArticleID = param
}

// FieldAuthorID is a field or column in the table Comment.
func (c *Comment) FieldAuthorID() string {
	return "authorID"
}

// SetAuthorID is a setter for the field or column AuthorID in the table Comment.
func (c *Comment) SetAuthorID(param int) {
	c.AuthorID = param
}

// FieldBody is a field or column in the table Comment.
func (c *Comment) FieldBody() string {
	return "body"
}

// SetBody is a setter for the field or column Body in the table Comment.
func (c *Comment) SetBody(param string) {
	c.Body = param
}

// AllField is a function to get all field or column in the table Comment.
func (c *Comment) AllField() (str []string) {
	str = []string{ 
		`createdAt`,
		`updatedAt`,
		`id`,
		`articleID`,
		`authorID`,
		`body`,
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
		case c.FieldAuthorID():
			values = append(values, c.AuthorID)
		case c.FieldBody():
			values = append(values, c.Body)
		case c.FieldCreatedAt():
			values = append(values, c.CreatedAt)
		case c.FieldUpdatedAt():
			values = append(values, c.UpdatedAt)
		case c.FieldID():
			values = append(values, c.ID)
		case c.FieldArticleID():
			values = append(values, c.ArticleID)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Comment.
func (c *Comment) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case c.FieldUpdatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field UpdatedAt")
			}
			c.SetUpdatedAt(val)
		case c.FieldID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type int. field ID")
			}
			c.SetID(val)
		case c.FieldArticleID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type int. field ArticleID")
			}
			c.SetArticleID(val)
		case c.FieldAuthorID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type int. field AuthorID")
			}
			c.SetAuthorID(val)
		case c.FieldBody():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Body")
			}
			c.SetBody(val)
		case c.FieldCreatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedAt")
			}
			c.SetCreatedAt(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

