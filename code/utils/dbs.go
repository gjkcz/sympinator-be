package utils

import (
	"os"
	"log"
	"fmt"
	"io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
/*// {{{ QuerySQLConn
/*
* calls callback upon successfully connecting to our pre-selected sql database
* expects function returning sql.Result, error
*/

func QuerySQLConn(dbname string, query string,queryVars ...interface{}) (*sql.Rows, error) {
	hostname := os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")
	// TODO: figure out how to securely connect to my own created database (no root/root password)
	db, err := sql.Open("mysql", "root:root@tcp("+hostname+")/"+dbname)

	if (err != nil) {
		log.Println("Could not connect to specified database; Error: ",err)
		db.Close()
		return nil, err
	}
	resultRows, err := db.Query(query, queryVars...)

	if (err != nil) {
		log.Println("Error while executing query: ",err)
		db.Close()
		return nil, err
	}

	defer db.Close() // TODO: apparently there is no reason to keep closing and re-opening
	return resultRows, err
}
// }}}

// PrintQueryResult(writer, rows *sql.Rows) {{{ helper func to dump mysql query results
func PrintQueryResult(w io.Writer,rows *sql.Rows) {

	// Stolen for convenience from stackoverflow answerhttps://stackoverflow.com/a/14500756
	// this is for testing purposes and will not appear in the final project
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}
	for rows.Next() {
		fmt.Println("printing row");
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Fprintf(w,"%#v\n", result)
	}
}
// }}}
