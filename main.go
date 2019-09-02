package main

import (
	dbp "github.com/celest1al/blog-auth-system/db"
	"github.com/celest1al/blog-auth-system/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setUpDB() {
	db := dbp.New()
	dbp.SetUpDB(db)
}

func main() {
	setUpDB()
	routes.SetupRoutes()
}
