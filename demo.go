package main

import (
	"github.com/doytowin/goooqo"
	"github.com/doytowin/goooqo/rdb"
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
	goooqo.BuildRestService[UserEntity, UserQuery]("/user/", userDataAccess)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
