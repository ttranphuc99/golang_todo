package dtos

import "todoapi/models"

type LoginDTO struct {
	Token string                `json:"token"`
	User  models.UserAccountDTO `json:"user"`
}
