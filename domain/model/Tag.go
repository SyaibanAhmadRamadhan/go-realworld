package model

type Tag struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}
