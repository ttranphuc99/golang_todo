package models

type Todo struct {
	ID          int64  `json:id,omitempty`
	Title       string `json:title`
	Content     string `json:content`
	Status      int    `json:status`
	OwnerId     string `json:ownerId,omitempty`
	CreatedTime string `json:createdTime,omitempty`
	UpdateTime  string `json:updatedTime,omitempty`
}
