package main

import (
	"github.com/doytowin/goooqo/core"
	"github.com/doytowin/goooqo/rdb"
	"github.com/doytowin/goooqo/web"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	buildEchoService[UserEntity, UserQuery](e, "/user/", userDataAccess)
	e.Logger.Fatal(e.Start(":9090"))
}

func buildEchoService[E core.Entity, Q core.Query](e *echo.Echo,
	prefix string, dataAccess core.TxDataAccess[E], m ...echo.MiddlewareFunc) {
	idPrefix := prefix + ":id"
	e.GET(idPrefix, func(c echo.Context) error {
		id := c.Param("id")
		entity, err := dataAccess.Get(c.Request().Context(), id)
		return writeResult(c, err, entity)
	}, m...)
	e.GET(prefix, func(c echo.Context) error {
		query := *new(Q)
		web.ResolveQuery(c.QueryParams(), &query)
		data, err := dataAccess.Page(c.Request().Context(), query)
		return writeResult(c, err, data)
	}, m...)
	e.PATCH(idPrefix, func(c echo.Context) error {
		entity := *new(E)
		if err := c.Bind(&entity); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		entity.SetId(&entity, c.Param("id"))
		data, err := dataAccess.Patch(c.Request().Context(), entity)
		return writeResult(c, err, data)
	}, m...)
	e.PATCH(prefix, func(c echo.Context) error {
		query := *new(Q)
		web.ResolveQuery(c.QueryParams(), &query)
		entity := *new(E)
		if err := c.Bind(&entity); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		data, err := dataAccess.PatchByQuery(c.Request().Context(), entity, query)
		return writeResult(c, err, data)
	}, m...)
	e.POST(prefix, func(c echo.Context) error {
		var entities []E
		if err := c.Bind(&entities); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		data, err := dataAccess.CreateMulti(c.Request().Context(), entities)
		return writeResult(c, err, data)
	}, m...)
	e.DELETE(idPrefix, func(c echo.Context) error {
		id := c.Param("id")
		data, err := dataAccess.Delete(c.Request().Context(), id)
		return writeResult(c, err, data)
	}, m...)
}

func writeResult(c echo.Context, err error, data any) error {
	response := core.Response{Data: data, Success: core.NoError(err), Error: core.ReadError(err)}
	return c.JSON(0, response)
}
