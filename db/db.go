package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func open() (*sql.DB, error) {

	wd, err := os.Getwd()
	path := wd + "/db/database.db"
	database, err := sql.Open("sqlite3", path)

	tableQuery := "create table if not exists rsa_keys(" +
		"id integer primary key autoincrement," +
		"name varchar(64)," +
		"key text);"

	_, err = database.Exec(tableQuery)
	if err != nil {
		fmt.Println("Creating table", err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func SaveRSAKey(name string, key string) {
	db, err := open()
	if err != nil {
		fmt.Println("Opening: ", err)
		return
	}

	query := "insert into rsa_keys (name, key) values (?,?);"
	_, err = db.Exec(query, name, key)
	if err != nil {
		fmt.Println("save: ", err)
		return
	}
}

func GetRSAKeys() {
	db, err := open()
	if err != nil {
		fmt.Println("Opening: ", err)
		return
	}

	query := "select id, name from rsa_keys;"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	fmt.Println("ID\tName")
	for rows.Next() {
		var id int
		var name string

		// Scan the values from the current row into variables
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		// Print the values as a table row
		fmt.Printf("%d\t%s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}

func GetRSAKey(id int) (string, error) {
	db, err := open()
	if err != nil {
		fmt.Println("Opening: ", err)
		return "", err
	}

	var key string

	query := "select key from rsa_keys where id=?;"
	row := db.QueryRow(query, id)

	err = row.Scan(&key)
	if err != nil {
		fmt.Println("Error querying for key: ", err)
		return "", err
	}

	if err := row.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return "", err
	}

	return key, nil
}

func DeleteRSAKey(id int) {
	db, err := open()
	if err != nil {
		fmt.Println("Opening: ", err)
		return
	}
	query := "delete from rsa_keys where id=?;"
	_, err = db.Exec(query, id)
	if err != nil {
		return
	}
}
