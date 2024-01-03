package lib

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/mattn/go-sqlite3"
)

func SqlError(err error, lokkingForFiels []string) (string, int) {
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		switch sqliteErr.Code {
		case sqlite3.ErrConstraint:
			return fmt.Sprintf("This %s already exists. Please use a different one", strings.Join(lokkingForFiels, " or ")), http.StatusConflict
		}
	}

	if err == sql.ErrNoRows {
		return fmt.Sprintf("%s not found.", strings.Join(lokkingForFiels, " or ")), http.StatusNotFound
	}

	return "Something went wrong if the problem persists contact us", http.StatusInternalServerError
}
