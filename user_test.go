package main

import (
	"context"
	"github.com/doytowin/goooqo/core"
	"github.com/doytowin/goooqo/rdb"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func Test(t *testing.T) {
	db := rdb.Connect("local.properties")
	defer rdb.Disconnect(db)
	tm := rdb.NewTransactionManager(db)
	ctx := context.Background()

	userDataAccess := rdb.NewTxDataAccess[UserEntity](tm)

	t.Run("Query", func(t *testing.T) {
		emailContain := "test"
		userQuery := UserQuery{EmailContain: &emailContain, MemoNull: core.PBool(true)}
		users, err := userDataAccess.Query(ctx, &userQuery)

		if err != nil {
			t.Error("Error", err)
		}
		if !(len(users) == 2 && users[0].Id == 2 &&
			users[0].Nickname == "test2" && users[0].Memo == nil) {
			t.Errorf("Data is not expected: %v", users)
		}
	})
}
