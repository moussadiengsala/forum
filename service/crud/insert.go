package crud

import (
	"database/sql"
	"fmt"
	"log"
)

func Insert(db *sql.DB, table string, columns []string, valuesToInsert ...interface{}) (int64, error) {
	fmt.Println(valuesToInsert...)
	// INSERT INTO id, firt_name, last_name VALUES ( ?, ?, ?)
	query := InsertQueryBuilder(table, columns...)

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err

	}
	defer stmt.Close()
	result, err := stmt.Exec(valuesToInsert...)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	log.Println("Successfully inserted in the database ✅")
	return id, nil
}
