package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/doytowin/goooqo/rdb"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func P[T any](t T) *T { return &t }

func Test(t *testing.T) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		defer db.Close()
	}

	tm := rdb.NewTransactionManager(db)
	ctx := context.Background()

	userDataAccess := rdb.NewTxDataAccess[UserEntity](tm)

	t.Run("Query", func(t *testing.T) {
		userQuery := UserQuery{EmailContain: P("test"), MemoNull: P(true)}
		users, err := userDataAccess.Query(ctx, &userQuery)
		//Executing SQL="SELECT id, username, email, nickname, memo, valid FROM t_user WHERE email LIKE ? AND memo IS NULL" args="[%test%]"

		if err != nil {
			t.Error("Error", err)
		}
		if !(len(users) == 2 && users[0].Id == 2 &&
			*users[0].Nickname == "test2" && users[0].Memo == nil) {
			t.Errorf("Data is not expected: %v", users)
		}
	})

	t.Run("OR query sample ", func(t *testing.T) {
		userQuery := UserQuery{UserOr: &[]UserQuery{
			{IdIn: &[]int{2, 4}, Valid: P(true)},
			{EmailContain: P("qq"), MemoNull: P(false)}}}
		users, _ := userDataAccess.Query(ctx, userQuery)
		//Executing SQL="SELECT id, username, email, nickname, memo, valid FROM t_user WHERE (id IN (?, ?) AND valid = ? OR email LIKE ? AND memo IS NOT NULL)" args="[2 4 true %qq%]"

		if !(len(users) == 2 && users[0].Id == 3 && users[1].Id == 4) {
			data, _ := json.Marshal(users)
			t.Errorf("Data is not expected: %v", string(data))
		}
	})
}
