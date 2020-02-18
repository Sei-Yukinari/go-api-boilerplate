package http

import (
	"api/infrastracture/db"
	"api/infrastracture/log"
	"api/infrastracture/router"
	"fmt"

	"github.com/fvbock/endless"

	"api/config"
	"api/infrastracture/env"

	_ "github.com/go-sql-driver/mysql"
)

func StartHttpServer() {
	c := config.Load()
	env.Load()
	sqlHandler, err := db.NewSQLHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}
	logger := log.NewLogger()
	if err := endless.ListenAndServe(":"+c.Server.Port, router.Handler(sqlHandler, logger)); err != nil {
		fmt.Printf("Error : [%s]", err)
	}
}
