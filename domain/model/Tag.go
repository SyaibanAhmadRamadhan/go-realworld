package model

type Tag struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}
