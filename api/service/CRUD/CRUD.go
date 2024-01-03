package service

import (
	"database/sql"
	"errors"
	"reflect"
)

// DBService defines the methods for common database operations.
type DBService interface {
	Create(query string, data ...interface{}) error
	SelectSingle(query string, condition []interface{}, dest ...interface{}) error
	SelectAllForPrimitive(query string, condition interface{}, dest interface{}) error
	SelectAllForStruct(query string, condition interface{}, dest interface{}, fields []string) error
	Update(query string, data ...interface{}) error
	Delete(query string, data ...interface{}) error
}

// DB implements the DBService interface.
type DB struct {
	instanceOfDB *sql.DB
}

func SqlService(db *sql.DB) DBService {
	return &DB{
		instanceOfDB: db,
	}
}

func (db *DB) Create(query string, data ...interface{}) error {
	_, err := db.instanceOfDB.Exec(query, data...)
	return err
}

func (db *DB) SelectSingle(query string, condition []interface{}, dest ...interface{}) error {
	err := db.instanceOfDB.QueryRow(query, condition...).Scan(dest...)
	return err
}

func (db *DB) SelectAllForPrimitive(query string, condition interface{}, dest interface{}) error {
	rows, err := db.instanceOfDB.Query(query, condition)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Ensure dest is a slice
	sliceValue := reflect.ValueOf(dest)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be a pointer to a slice")
	}

	sliceElemType := sliceValue.Elem().Type().Elem()

	for rows.Next() {
		// Create a new instance of the slice element type
		elem := reflect.New(sliceElemType).Interface()
		err = rows.Scan(elem)
		if err != nil {
			return err
		}

		// Append the scanned element to the destination slice
		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), reflect.ValueOf(elem).Elem()))
	}

	return nil
}

func (db *DB) SelectAllForStruct(query string, condition interface{}, dest interface{}, fields []string) error {
	rows, err := db.instanceOfDB.Query(query, condition)
	if err != nil {
		return err
	}
	defer rows.Close()

	typeDest := reflect.TypeOf(dest)
	entries := reflect.New(typeDest.Elem()).Elem()
	entrie := reflect.New(typeDest.Elem().Elem()).Elem()

	// fmt.Println(entrie.Type().Field(0).Name, entries.Type())

	for rows.Next() {
		// Scan the values into the specified fields
		values := make([]interface{}, len(fields))
		for i, fieldName := range fields {
			fieldValue := entrie.FieldByName(fieldName)
			values[i] = fieldValue.Addr().Interface()
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}

		// Append the entrie to the entries slice
		entries = reflect.Append(entries, entrie)
	}

	// Set the result to the original destination
	reflect.ValueOf(dest).Elem().Set(entries)

	return nil
}

func (db *DB) Update(query string, data ...interface{}) error {
	_, err := db.instanceOfDB.Exec(query, data...)
	return err
}

func (db *DB) Delete(query string, data ...interface{}) error {
	_, err := db.instanceOfDB.Exec(query, data...)
	return err
}
