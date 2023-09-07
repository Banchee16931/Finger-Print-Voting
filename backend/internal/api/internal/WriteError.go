package internal

import (
	"encoding/json"
	"finger-print-voting-backend/internal/cerr"
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, msg error, metadata ...interface{}) {
	log.Printf("Returning with status code: %d\n", status)
	log.Printf("Returned Error Message: %s\n", msg)

	w.WriteHeader(status)
	if len(metadata) == 0 {
		json.NewEncoder(w).Encode(&cerr.CommonError{
			Message: msg.Error(),
		})
	} else if len(metadata) == 1 {
		json.NewEncoder(w).Encode(&cerr.CommonError{
			Message:  msg.Error(),
			Metadata: metadata[0],
		})
	} else {
		json.NewEncoder(w).Encode(&cerr.CommonError{
			Message:  msg.Error(),
			Metadata: metadata,
		})
	}
}
