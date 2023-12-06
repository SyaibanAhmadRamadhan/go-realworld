package model

type UserFollow struct {
	UserID     int `bson:"userID"`
	FollowedID int `bson:"followedID"`
}
