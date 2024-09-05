package db

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) error {
	createTablePostSql := `CREATE TABLE IF NOT EXISTS posts (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	title VARCHAR(200) NOT NULL,
    	description TEXT NOT NULL,
    	tags VARCHAR (150) NOT NULL,
    	user VARCHAR(50) NOT NULL,
    	status INTEGER NOT NULL,
    	create_at TIMESTAMP NOT NULL
	)`

	createAnswerResponseSql := `CREATE TABLE IF NOT EXISTS answers (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	post_id INTEGER NOT NULL,
    	response TEXT NOT NULL,
    	user VARCHAR(50) NOT NULL,
    	create_at TIMESTAMP NOT NULL
	)`

	createTableTagsSql := `CREATE TABLE IF NOT EXISTS tags (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	tags VARCHAR(15) NOT NULL UNIQUE
	)`

	createTablePostTagsSql := `CREATE TABLE IF NOT EXISTS tags (
    	post_id INTEGER NOT NULL,
    	tag_id INTEGER NOT NULL
	)`

	createTableUsersSql := `CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name VARCHAR(150) NOT NULL,
    	login VARCHAR(50) NOT NULL,
    	email VARCHAR(150) NOT NULL,
    	password VARCHAR(500) NOT NULL,
    	create_at TIMESTAMP NOT NULL
	)`

	ddl := []string{createTablePostSql,
		createAnswerResponseSql,
		createTableTagsSql,
		createTablePostTagsSql,
		createTableUsersSql}

	for _, sql := range ddl {
		_, err := db.Exec(sql)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
