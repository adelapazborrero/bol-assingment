package main

import (
	"github.com/Abeldlp/bol-assignment/user-api/config"
	"github.com/Abeldlp/bol-assignment/user-api/migrations"
	"github.com/Abeldlp/bol-assignment/user-api/route"
)

func main() {
	//Initial setup
	config.InitializeEnvironmentVariables()
	config.InitializeDatabase()

	//Migrations
	migrations.Migrate(config.DB)

	r := config.InitializeServer()

	//Routes
	route.AppendUserRoute(r)

	r.Run(":8000")
}
