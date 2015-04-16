package helpers

import (
	coreglobals "collexy/core/globals"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func CheckIfDbInstalled() (isInstalled bool) {
	db := coreglobals.Db
	var res sql.NullString

	query := fmt.Sprintf("SELECT to_regclass('%s.user');", "public")
	err := db.QueryRow(query).Scan(&res)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No tables with that schema_name.")
	case err != nil:
		log.Fatal(err)
	default:
		if res.Valid {
			isInstalled = true
		} else {
			isInstalled = false
		}
	}

	log.Println("isinstalled:::::")
	log.Println(isInstalled)
	return
}
