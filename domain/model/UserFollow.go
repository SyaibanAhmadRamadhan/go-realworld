package model

type UserFollow struct {
	UserID     string   `bson:"userID"`
	FollowedID []string `bson:"followedID"`
}
