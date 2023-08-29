package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"net/http"
)

func (srv *Server) HandleUser(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(types.UserContext)

	user, ok := userCtx.(types.UserDetails)
	if !ok {
		WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate JWT"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
