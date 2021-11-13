package models

type Todo struct {
	ID      int64  `json:id`
	Content string `json:content`
	Status  int    `json:status`
}
