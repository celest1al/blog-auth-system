package db

import (
	"fmt"

	"github.com/celest1al/blog-auth-system/models"
	"github.com/celest1al/blog-auth-system/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // register postgres driver
)

// These would actually come from enviroment variables.
// An example of that is available here:
const (
	user    = "postgres"
	pwd     = ""
	url     = "127.0.0.1"
	port    = "5432"
	dbName  = "mydb"
	sslmode = "disable"
)

// SetUpDB Resets database and migrate table users schema
func SetUpDB(db *gorm.DB) {
	fmt.Println("Resetting database")
	db.DropTableIfExists(&models.User{})
	db.AutoMigrate(&models.User{})
}

// New creates a database connection
func New() *gorm.DB {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pwd, url, port, dbName, sslmode)
	db, err := gorm.Open("postgres", dbConnString)
	utils.Must(err)
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	err = db.DB().Ping()
	utils.Must(err)
	return db
}
