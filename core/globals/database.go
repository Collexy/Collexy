package globals

import(
	"database/sql"
    _ "github.com/lib/pq"
	"log"
	"collexy/globals"
	"fmt"
)
var Db *sql.DB

func SetupDB() (db *sql.DB) {
	connString := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s", globals.Conf.DbName, globals.Conf.DbUser, globals.Conf.DbPassword, globals.Conf.SslMode)
	//log.Println(connString)
	db, err := sql.Open(globals.Conf.DbUser, connString)
	//db, err := sql.Open("postgres", "dbname=collexy user=postgres password=hathat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

//var Db = helpers.SetupDB()