package main

import (
	"github.com/doytowin/goooqo"
)

type UserEntity struct {
	goooqo.Int64Id
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	Memo     *string `json:"memo,omitempty"`
	Valid    *bool   `json:"valid,omitempty"`
}

func (u UserEntity) GetTableName() string {
	return "t_user"
}

type UserQuery struct {
	goooqo.PageQuery
	IdGt         *int
	IdIn         *[]int
	EmailContain *string
	MemoNull     *bool
	Valid        *bool
	UserOr       *[]UserQuery
}
