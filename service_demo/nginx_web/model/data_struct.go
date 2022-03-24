package model

type PostContent struct {
	UserID       string `json:"user_id"`
	TimeStamp    string `json:"time_stamp"`
	TextContent  string `json:"text_content"`
	MediaContent string `json:"media_content"`
}

type Media struct {
	UserID       string `json:"user_id"`
	TimeStamp    string `json:"time_stamp"`
	MediaContent string `json:"media_content"`
}

type Text struct {
	UserID      string `json:"user_id"`
	TimeStamp   string `json:"time_stamp"`
	TextContent string `json:"text_content"`
}

type ChError struct {
	IsError  bool
	ErrorMsg string
}
