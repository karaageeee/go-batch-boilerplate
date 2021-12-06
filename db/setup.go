package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// Setup might call only once
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/karaageeee/go-batch-boilerplate/config"
)

var dbConnection *gorm.DB

// Setup is to set DB connection to singleton instance
func Setup() {
	log.Info("DB setup start")
	db, err := gorm.Open("postgres", connectionInfo())
	if err != nil {
		log.Fatal(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	dbConnection = db
	log.Info("DB setup done")
}

// GetDBConnection is to get db connection instance
func GetDBConnection() *gorm.DB {
	if dbConnection == nil {
		Setup()
	}
	return dbConnection
}

// ConnectionInfo returns string for gorm.Open
func connectionInfo() string {
	env := os.Getenv("ENV")
	sslmode := "disable"
	if env == "production" || env == "staging" {
		sslmode = "require"
	}
	connInfo := config.GetDBConnectionInfo()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", connInfo.HOST, connInfo.PORT, connInfo.User, connInfo.Pass, connInfo.Name, sslmode)
}
