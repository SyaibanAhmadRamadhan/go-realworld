package model

type User struct {
	ID       string  `bson:"_id"`
	Email    string  `bson:"email"`
	Username string  `bson:"username"`
	Password string  `bson:"password"`
	Image    string  `bson:"image"` // 'https://realworld-temp-api.herokuapp.com/images/smiley-cyrus.jpeg'
	Bio      *string `bson:"bio"`
	Demo     bool    `bson:"demo"`
}

// -- CreateTable
// CREATE TABLE "_UserFavorites" (
// "A" INTEGER NOT NULL,
// "B" INTEGER NOT NULL
// );
//
// -- CreateTable
// CREATE TABLE "_UserFollows" (
// "A" INTEGER NOT NULL,
// "B" INTEGER NOT NULL
// );
