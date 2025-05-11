package main

import (
	"github.com/doytowin/goooqo/core"
	"github.com/doytowin/goooqo/rdb"
	"github.com/doytowin/goooqo/web"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.DebugLevel)

	db := rdb.Connect("local.properties")
	defer rdb.Disconnect(db)
	tm := rdb.NewTransactionManager(db)

	// Init data layer
	core.RegisterJoinTable("role", "user", "a_user_and_role")
	UserDataAccess = rdb.NewTxDataAccess[UserEntity](tm)
	RoleDataAccess = rdb.NewTxDataAccess[RoleEntity](tm)

	// Init web layer
	web.BuildRestService[UserEntity, UserQuery]("/user/", UserDataAccess)
	web.BuildRestService[RoleEntity, RoleQuery]("/role/", RoleDataAccess)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
