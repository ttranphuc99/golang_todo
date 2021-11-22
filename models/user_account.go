package models

import "database/sql"

type UserAccount struct {
	LoginId     string         `json:loginId`
	Fullname    string         `json:fullname`
	Password    string         `json:password`
	Role        int            `json:role`
	Status      int            `json:status`
	CreatedTime string         `json:createdTime`
	UpdatedTime string         `json:updatedTime`
	UpdatedBy   sql.NullString `json:updatedBy`
}

type UserAccountDTO struct {
	LoginId     string `json:loginId`
	Fullname    string `json:fullname`
	Role        int    `json:role`
	Status      int    `json:status`
	CreatedTime string `json:createdTime`
	UpdatedTime string `json:updatedTime`
	UpdatedBy   string `json:updatedBy`
}
