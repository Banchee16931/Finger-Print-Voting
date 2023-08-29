package api

import (
	"encoding/json"
	"errors"
	"finger-print-voting-backend/internal/api/auth"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (srv *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var userReq types.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := userReq.Validate(); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := srv.db.GetUser(userReq.Username)
	if err != nil {
		if errors.Is(err, cerr.ErrNotFound) {
			WriteError(w, http.StatusUnauthorized, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		WriteError(w, http.StatusUnauthorized, fmt.Errorf("incorrect password"))
		return
	}

	log.Println("Correct password")

	jwt, err := auth.GenerateJWT(srv.passwordSecret, user.Username)
	if err != nil {
		log.Println("Error: ", err.Error())
		WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate JWT"))
		return
	}

	log.Println("jwt: ", jwt)

	w.Header().Add("Set-Cookie", fmt.Sprintf("Autherization: Bearer {%s}", jwt))
	w.WriteHeader(http.StatusOK)
}
