package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var RunMode string
var ServerPort string
var MongoServer string
var MongoPort string
var MongoUsername string
var MongoPassword string

var DatabaseConnectionString string
var DatabaseName string

func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}

	log.Println("RUN MODE:", RunMode)

	if RunMode != PRODUCTION {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}

	ServerPort = os.Getenv("SERVER_PORT")

	MongoServer = os.Getenv("MONGO_SERVER")
	MongoPort = os.Getenv("MONGO_PORT")
	MongoUsername = os.Getenv("MONGO_USERNAME")
	MongoPassword = os.Getenv("MONGO_PASSWORD")

	DatabaseConnectionString = "mongodb://" + MongoUsername + ":" + MongoPassword + "@" + MongoServer + ":" + MongoPort
	DatabaseName = os.Getenv("DATABASE_NAME")
}
