package models

import "time"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Tweet struct {
	UserName   string    `json:"username"`
	Tweet      string    `json:"tweet"`
	CreateTime time.Time `json:"create_time"`
}

type GetAllTweets struct {
	GetTweet []GetTweet `json:"tweets"`
}

type GetTweet struct {
	Id           int64  `json:"id"`
	Tweet        string `json:"tweet"`
	CreatedAt    string `json:"created_at"`
	LikeCount    int64  `json:"like_count"`
	RetweetCount int64  `json:"retweet_count"`
}
