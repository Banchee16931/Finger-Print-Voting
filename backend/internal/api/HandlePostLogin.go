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

func (srv *Server) HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	var userReq types.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	if err := userReq.Validate(); HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	user, err := srv.db.GetUser(userReq.Username)
	if err != nil {
		if errors.Is(err, cerr.ErrNotFound) {
			HTTPError(w, http.StatusUnauthorized, err)
			return
		}
		HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if HTTPError(w, http.StatusUnauthorized, err) {
		return
	}

	log.Println("Correct password")

	jwt, err := auth.GenerateJWT(srv.passwordSecret, user.Username)
	if err != nil {
		log.Println("JWT Error: ", err.Error())
		HTTPError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate JWT"))
		return
	}

	log.Println("jwt: ", jwt)

	w.Header().Add("Set-Cookie", fmt.Sprintf("Autherization: Bearer {%s}", jwt))
	w.WriteHeader(http.StatusOK)
}
