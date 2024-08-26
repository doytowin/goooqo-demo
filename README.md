[![License](https://img.shields.io/:license-apache-brightgreen.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

Demo for GoooQo
---

## Introduction

This repo shows how to use [GoooQo](https://github.com/doytowin/goooqo) to access a single table.

The [test.db](test.db) is a database of sqlite which contains a simple table `t_user` with 4 rows.
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
  - http://localhost:9090/user/?emailContain=qq
  - http://localhost:9090/user/?emailContain=qq&IdGt=2
  - http://localhost:9090/user/?emailContain=qq&memoNull=true
  - http://localhost:9090/user/?emailContain=qq&memoNull=false

And you can also try the `POST/PUT/PATCH/DELETE` operations listed in [user.http](user.http).

> NOTE: The web module of GoooQo is for tests and demos only, not a core feature.

License
---
This project is under the [Apache Licence v2](https://www.apache.org/licenses/LICENSE-2.0).
