package models

type Todo struct {
	ID          int64  `json:id`
	Title       string `json:title`
	Content     string `json:content`
	Status      int    `json:status`
	OwnerId     string `json:ownerId`
	CreatedTime string `json:createdTime`
	UpdateTime  string `json:updatedTime`
}
