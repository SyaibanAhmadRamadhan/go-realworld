package model

type UserFollow struct {
	UserId     string   `bson:"userId"`
	FollowedId []string `bson:"followedId"`
}
