package model

type User struct {
	Id       string  `bson:"_id"`
	Email    string  `bson:"email"`
	Username string  `bson:"username"`
	Password string  `bson:"password"`
	Image    string  `bson:"image"` // 'https://realworld-temp-api.herokuapp.com/images/smiley-cyrus.jpeg'
	Bio      *string `bson:"bio"`
	Demo     bool    `bson:"demo"`
}
