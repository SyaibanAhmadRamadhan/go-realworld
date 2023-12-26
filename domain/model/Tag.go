package model

type Tag struct {
	Id   string `bson:"_id"    order:"true"`
	Name string `bson:"name"   order:"true"`
}
