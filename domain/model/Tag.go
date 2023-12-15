package model

type Tag struct {
	ID   string `bson:"_id"    order:"true"`
	Name string `bson:"name"   order:"true"`
}
