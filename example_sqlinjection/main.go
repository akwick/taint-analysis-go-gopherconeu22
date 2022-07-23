// CWE-89 Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')
package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3" // init() only
)

func main() {
	db := openDatabase()
	defer db.Close()
	a := Authentication{"Gopher", source()}

	query(a.Password, a, db) // query is a sink
}

func openDatabase() *sql.DB {
	const database = "/tmp/db_test.db"
	os.Remove(database)
	db, _ := sql.Open("sqlite3", database)
	return db
}

func source() string {
	return "My Gopher Secret"
}
