package models

type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cast        string `json:"cast"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
}
