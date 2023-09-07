package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"net/http"
)

// Gets the details about the user accessing this endpoint
func (srv *Server) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(types.UserContext) // get the user data from the auth middleware

	user, ok := userCtx.(types.UserDetails) // check it is valid
	if !ok {
		HTTPError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate JWT"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user) // output the user
}
