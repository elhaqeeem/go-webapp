package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB SERVER (AWS)
func dbvar() string {

	// THIS IS FOT GET THE ENV VARIABLES IN LOCAL, IN HEROKU IT IS DIFFERENT
	//err := godotenv.Load(".env")
	//if err != nil {
	//	fmt.Println("Error loading .env file / HEROKU DEPLOY ?") // IF NO .env -> IS RUNNING ON HEROKU
	//}

	// CONFIG VARS
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	SSL_MODE := os.Getenv("SSL_MODE")
	//sslrootcert := os.Getenv("sslrootcert")
	//sslkey := os.Getenv("sslkey")
	//sslcert := os.Getenv("sslcert")

	dsn := ("host=" + DB_HOST +
		" user=" + DB_USER +
		" password=" + DB_PASSWORD +
		" dbname=" + DB_NAME +
		" port=" + DB_PORT +
		" sslmode=" + SSL_MODE)
	//" sslrootcert=" + sslrootcert +
	//" sslkey=" + sslkey +
	//" sslcert=" + sslcert
	//

	return dsn
}

var DB = func() (db *gorm.DB) {

	if db, err := gorm.Open(postgres.Open(dbvar()), &gorm.Config{}); err != nil {
		fmt.Println("Connection to database failed", err)
		panic(err)
	} else {
		fmt.Println("Connected to database")
		return db
	}
}()
