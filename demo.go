package main

import (
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

	userDataAccess := rdb.NewTxDataAccess[UserEntity](tm)
	web.BuildRestService[UserEntity, UserQuery]("/user/", userDataAccess)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
