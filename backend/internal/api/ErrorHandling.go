package api

import (
	"database/sql"
	"finger-print-voting-backend/internal/api/internal"
	"fmt"
	"net/http"
)

// Will write an error to the ResponseWriter, with the given status, if it exists and return true,
// otherwise will return false
func HTTPError(w http.ResponseWriter, status int, err error, metadata ...interface{}) bool {
	if err == nil {
		return false
	}

	internal.WriteError(w, status, err, metadata...)
	return true
}

// Will write an error to the ResponseWriter, with the given status, if it exists and return true,
// otherwise will return false
//
// Also handles the rollback procedure on the transaction and any error handling as a result of that
func HTTPErrorWithRollback(tx *sql.Tx, w http.ResponseWriter, status int, err error, metadata ...interface{}) bool {
	if err == nil {
		return false
	}

	rollErr := tx.Rollback()
	if rollErr != nil {
		internal.WriteError(w, http.StatusInternalServerError, fmt.Errorf("%s: %w", rollErr.Error(), err), metadata...)
		return true
	}

	internal.WriteError(w, http.StatusInternalServerError, err, metadata...)
	return true
}
