package dtos

type TodoDTO struct {
	ID          int64  `json:"id",omitempty`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Status      int16  `json:"status"`
	OwnerId     string `json:"ownerId",omitempty`
	CreatedTime string `json:"createdTime",omitempty`
	UpdateTime  string `json:"updatedTime",omitempty`
}
