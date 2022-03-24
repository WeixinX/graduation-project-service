package db

type Media struct {
	UserID       string `json:"user_id"`
	TimeStamp    string `json:"time_stamp"`
	MediaContent string `json:"media_content"`
}
