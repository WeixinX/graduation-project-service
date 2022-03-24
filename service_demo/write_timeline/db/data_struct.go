package db

type Text struct {
	UserID    string `json:"user_id"`
	TimeStamp string `json:"time_stamp"`
	Content   string `json:"content"`
}