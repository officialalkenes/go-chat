package model

type Chat Struct {
	ID string `json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}