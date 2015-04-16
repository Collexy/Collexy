package globals

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func SetupDB() (db *sql.DB) {

	connString := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s", Conf.DbName, Conf.DbUser, Conf.DbPassword, Conf.SslMode)
	//log.Println(connString)
	db, err := sql.Open(Conf.DbUser, connString)
	//db, err := sql.Open("postgres", "dbname=collexy user=postgres password=hathat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

//var Db = helpers.SetupDB()
