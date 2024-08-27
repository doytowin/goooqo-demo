[![License](https://img.shields.io/:license-apache-brightgreen.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

Demo for GoooQo
---

## Introduction

This repo shows how to use [GoooQo](https://github.com/doytowin/goooqo) to access a single table.

The [test.db](test.db) is a database of sqlite which contains a simple table `t_user` with 4 rows.

| id | username | password | email        | mobile      | nickname | memo | valid |
|----|----------|----------|--------------|-------------|----------|------|-------|
| 1  | f0rb     | 123456   | f0rb@163.com | 18888888881 | test1    |      | true  |
| 2  | user2    | 123456   | test2@qq.com | 18888888882 | test2    |      | false |
| 3  | user3    | 123456   | test3@qq.com | 18888888883 | test3    | memo | true  |
| 4  | user4    | 123456   | test4@qq.com | 18888888884 | test4    |      | true  |

With GoooQo, we only need to define the structs `UserEntity` and `UserQuery` in [user.go](user.go),
where `UserEntity` corresponds to the table `t_user` and `UserQuery` is used to build the query clause.
The `goooqo.PageQuery` embedded in `UserQuery` helps to implement sorting and paging operations.

Then we build the web APIs listened on `http://localhost:9090/user/` by the following two lines:

```go
userDataAccess := rdb.NewTxDataAccess[UserEntity](tm)
goooqo.BuildRestService[UserEntity, UserQuery]("/user/", userDataAccess)
```

Run the main method in [demo.go](demo.go) and visit the the following URLs to check the response:
  - http://localhost:9090/user/
  - http://localhost:9090/user/3
  - http://localhost:9090/user/?pageNumber=2&pageSize=2
  - http://localhost:9090/user/?sort=id,desc
  - http://localhost:9090/user/?sort=memo,desc%3Bemail,desc
  - http://localhost:9090/user/?idIn=1,3
  - http://localhost:9090/user/?idGt=2&emailContain=qq
  - http://localhost:9090/user/?emailContain=qq
  - http://localhost:9090/user/?emailContain=qq&memoNull=true
  - http://localhost:9090/user/?emailContain=qq&memoNull=false

And there are also some sample `POST/PUT/PATCH/DELETE` operations listed in [user.http](user.http) to take a try.

> NOTE:The web module of GoooQo is for tests and demos only, not a core feature.

License
---
This project is under the [Apache Licence v2](https://www.apache.org/licenses/LICENSE-2.0).
