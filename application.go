package main

import (
	"restaurentmanagement/config"
	"restaurentmanagement/pkg"
	"restaurentmanagement/pkg/db"
)

func main() {
	srv := config.New()
	db.GetDmManager()
	pkg.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
