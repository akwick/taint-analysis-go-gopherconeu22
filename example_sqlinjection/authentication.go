package main

import "database/sql"

type Authentication struct {
	Username string
	Password string `datapolicy:"userinput"`
}

func query(q string, a Authentication, db *sql.DB) { // sink
	_, _ = db.Query("SELECT * FROM users WHERE id =" + q)
	_, _ = db.Query("SELECT * FROM users WHERE id =" + a.Password)
}
