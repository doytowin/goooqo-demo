[![License](https://img.shields.io/:license-apache-brightgreen.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

Demo for GoooQo
---

## Introduction

This repo shows how to use [GoooQo](https://github.com/doytowin/goooqo) to access a single table.

The [test.db](test.db) is a database of sqlite which contains a simple table `t_user` with 4 rows.
With GoooQo, we just defined the structs `UserEntity` and `UserQuery` in [user.go](user.go),
and build a set of web APIs listened on `http://localhost:9090/user/` by the following code.
The `UserQuery` embeds `goooqo.PageQuery` so it has paging capabilities.

```go
userDataAccess := rdb.NewTxDataAccess[UserEntity](tm)
goooqo.BuildRestService[UserEntity, UserQuery]("/user/", userDataAccess)
```

Run the main method in [demo.go](demo.go) and visit the the following urls to check the response:
  - http://localhost:9090/user/
  - http://localhost:9090/user/3
  - http://localhost:9090/user/?PageNumber=2&PageSize=2
  - http://localhost:9090/user/?Sort=id,desc
  - http://localhost:9090/user/?Sort=Memo,desc%3Bemail,desc
  - http://localhost:9090/user/?EmailContain=qq
  - http://localhost:9090/user/?EmailContain=qq&IdGt=2
  - http://localhost:9090/user/?EmailContain=qq&MemoNull=true
  - http://localhost:9090/user/?EmailContain=qq&MemoNull=false

And you can also visit http://localhost:9090/user/ by `POST/PUT/PATCH/DELETE` methods.

>NOTE: The web module of GoooQo is for tests and demo only, not core feature.

License
---
This project is under the [Apache Licence v2](https://www.apache.org/licenses/LICENSE-2.0).
