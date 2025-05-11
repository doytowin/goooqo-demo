package main

import (
	. "github.com/doytowin/goooqo/core"
)

//go:generate gooogen

type RoleEntity struct {
	IntId
	RoleName     *string
	RoleCode     *string
	CreateUserId *int

	Users []UserEntity `entitypath:"role,user" json:"users,omitempty"`
}

type RoleQuery struct {
	PageQuery
	IdIn          *[]int
	RoleNameStart *string
	Valid         *bool

	WithUsers *UserQuery
}

var RoleDataAccess TxDataAccess[RoleEntity]
