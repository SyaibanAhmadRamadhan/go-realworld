package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

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

// FieldDemo is a field or column in the table User.
func (u *User) FieldDemo() string {
	return "demo"
}

// SetDemo is a setter for the field or column Demo in the table User.
func (u *User) SetDemo(param bool) {
	u.Demo = param
}

// FieldID is a field or column in the table User.
func (u *User) FieldID() string {
	return "_id"
}

// SetID is a setter for the field or column ID in the table User.
func (u *User) SetID(param string) {
	u.ID = param
}

// FieldEmail is a field or column in the table User.
func (u *User) FieldEmail() string {
	return "email"
}

// SetEmail is a setter for the field or column Email in the table User.
func (u *User) SetEmail(param string) {
	u.Email = param
}

// FieldUsername is a field or column in the table User.
func (u *User) FieldUsername() string {
	return "username"
}

// SetUsername is a setter for the field or column Username in the table User.
func (u *User) SetUsername(param string) {
	u.Username = param
}

// FieldPassword is a field or column in the table User.
func (u *User) FieldPassword() string {
	return "password"
}

// SetPassword is a setter for the field or column Password in the table User.
func (u *User) SetPassword(param string) {
	u.Password = param
}

// FieldImage is a field or column in the table User.
func (u *User) FieldImage() string {
	return "image"
}

// SetImage is a setter for the field or column Image in the table User.
func (u *User) SetImage(param string) {
	u.Image = param
}

// FieldBio is a field or column in the table User.
func (u *User) FieldBio() string {
	return "bio"
}

// SetBio is a setter for the field or column Bio in the table User.
func (u *User) SetBio(param *string) {
	u.Bio = param
}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str []string) {
	str = []string{ 
		`_id`,
		`email`,
		`username`,
		`password`,
		`image`,
		`bio`,
		`demo`,
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
		case u.FieldID():
			values = append(values, u.ID)
		case u.FieldEmail():
			values = append(values, u.Email)
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
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct User.
func (u *User) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
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
		case u.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			u.SetID(val)
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
		case u.FieldPassword():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Password")
			}
			u.SetPassword(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

