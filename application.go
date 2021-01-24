package main

import (
	"restaurentManagement/config"
	"restaurentManagement/pkg"
	"restaurentManagement/pkg/db"
)

func main() {
	srv := config.New()
	db.GetDmManager()
	pkg.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
