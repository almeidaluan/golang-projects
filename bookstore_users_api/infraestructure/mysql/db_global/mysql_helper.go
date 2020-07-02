package db_global

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Client *sql.DB
)

func Init(){

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"))


	var err error
	Client, err := sql.Open("mysql",datasourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println(" Database  successfuly configured")
}
