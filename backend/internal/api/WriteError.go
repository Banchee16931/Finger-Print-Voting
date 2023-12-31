package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/cerr"
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, msg error) {
	log.Printf("Returning with status code: %d\n", status)
	log.Printf("Returned Error Message: %s\n", msg)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&cerr.CommonError{
		Message: msg.Error(),
	})
}
