package model

type Note struct {
	ID           int     `json:"id"`
	CategoryID   *int    `json:"category_id"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	CategoryName *string `json:"category_name"`
}
