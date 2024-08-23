package main

import (
	"github.com/doytowin/goooqo"
	"github.com/doytowin/goooqo/core"
)

type UserEntity struct {
	core.Int64Id
	Username string
	Email    string
	Mobile   string
	Nickname string
	Memo     *string
	Valid    bool
}

func (u UserEntity) GetTableName() string {
	return "t_user"
}

type UserQuery struct {
	goooqo.PageQuery
	IdGt         *int
	EmailContain *string
	MemoNull     *bool
}
