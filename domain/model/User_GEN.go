// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/ginfra/gdb/generator.go. 

package model

import (
	"errors"
)

// UserTableName this table or collection name
const UserTableName string = "user"

// NewUser is a struct with pointer that represents the table User in the database.
func NewUser() *User {
	return &User{}
}

// NewUserWithOutPtr is a struct without pointer that represents the table User in the database.
func NewUserWithOutPtr() User {
	return User{}
}

// FieldUsername is a field or column in the table User.
func (u *User) FieldUsername() string {
	return "username"
}

// SetUsername is a setter for the field or column Username in the table User.
func (u *User) SetUsername(param string) string {
	u.Username = param
	return "username"
}

// FieldPassword is a field or column in the table User.
func (u *User) FieldPassword() string {
	return "password"
}

// SetPassword is a setter for the field or column Password in the table User.
func (u *User) SetPassword(param string) string {
	u.Password = param
	return "password"
}

// FieldImage is a field or column in the table User.
func (u *User) FieldImage() string {
	return "image"
}

// SetImage is a setter for the field or column Image in the table User.
func (u *User) SetImage(param string) string {
	u.Image = param
	return "image"
}

// FieldBio is a field or column in the table User.
func (u *User) FieldBio() string {
	return "bio"
}

// SetBio is a setter for the field or column Bio in the table User.
func (u *User) SetBio(param *string) string {
	u.Bio = param
	return "bio"
}

// FieldDemo is a field or column in the table User.
func (u *User) FieldDemo() string {
	return "demo"
}

// SetDemo is a setter for the field or column Demo in the table User.
func (u *User) SetDemo(param bool) string {
	u.Demo = param
	return "demo"
}

// FieldId is a field or column in the table User.
func (u *User) FieldId() string {
	return "_id"
}

// SetId is a setter for the field or column Id in the table User.
func (u *User) SetId(param string) string {
	u.Id = param
	return "_id"
}

// FieldEmail is a field or column in the table User.
func (u *User) FieldEmail() string {
	return "email"
}

// SetEmail is a setter for the field or column Email in the table User.
func (u *User) SetEmail(param string) string {
	u.Email = param
	return "email"
}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str []string) {
	str = []string{ 
		`email`,
		`username`,
		`password`,
		`image`,
		`bio`,
		`demo`,
		`_id`,
	}
	return
}

// OrderFields is a function to get all field or column in the table User.
func (u *User) OrderFields() (str []string) {
	str = []string{ 
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table User.
func (u *User) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case u.FieldUsername():
			values = append(values, u.Username)
		case u.FieldPassword():
			values = append(values, u.Password)
		case u.FieldImage():
			values = append(values, u.Image)
		case u.FieldBio():
			values = append(values, u.Bio)
		case u.FieldDemo():
			values = append(values, u.Demo)
		case u.FieldId():
			values = append(values, u.Id)
		case u.FieldEmail():
			values = append(values, u.Email)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct User.
func (u *User) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case u.FieldPassword():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Password")
			}
			u.SetPassword(val)
		case u.FieldImage():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Image")
			}
			u.SetImage(val)
		case u.FieldBio():
			val, ok := value.(*string)
			if !ok {
				return errors.New("invalid type *string. field Bio")
			}
			u.SetBio(val)
		case u.FieldDemo():
			val, ok := value.(bool)
			if !ok {
				return errors.New("invalid type bool. field Demo")
			}
			u.SetDemo(val)
		case u.FieldId():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Id")
			}
			u.SetId(val)
		case u.FieldEmail():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Email")
			}
			u.SetEmail(val)
		case u.FieldUsername():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Username")
			}
			u.SetUsername(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

