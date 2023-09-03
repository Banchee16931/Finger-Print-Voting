package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (srv *Server) HandlePostRegistrationAcceptance(w http.ResponseWriter, r *http.Request) {
	var registrationReq types.UserAcceptanceRequest

	// get the request into a GoLang type
	err := json.NewDecoder(r.Body).Decode(&registrationReq)
	if HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	// get the registrant being accepted or declined
	registrant, err := srv.db.GetRegistrant(registrationReq.RegistrantID)
	if HTTPError(w, http.StatusNotFound, err) {
		return
	}

	// create transaction so all database edits are all or nothing in nature
	tx, err := srv.db.Begin()
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	// delete the registrant data from the database
	err = srv.db.DeleteRegistrant(tx, registrationReq.RegistrantID)
	if HTTPErrorWithRollback(tx, w, http.StatusInternalServerError, err) {
		return
	}

	// if accepted the new voter account needs to be created
	if registrationReq.Accepted {
		// check the request is valid
		err = registrationReq.Validate()
		if HTTPError(w, http.StatusBadRequest, err) {
			return
		}

		log.Println("user accepted: ", registrationReq.Username)

		// encrypt password
		excryptedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationReq.Password), bcrypt.MinCost)
		if HTTPErrorWithRollback(tx, w, http.StatusInternalServerError, err) {
			return
		}

		// transfer details to new voter account
		voter := types.Voter{
			User: types.User{
				Username:  registrationReq.Username,
				Password:  string(excryptedPassword),
				Admin:     false,
				FirstName: registrant.FirstName,
				LastName:  registrant.LastName,
			},
			PhoneNo:     registrant.PhoneNo,
			Email:       registrant.Email,
			Fingerprint: registrant.Fingerprint,
			Location:    registrant.Location,
		}

		// store new voter's details
		err = srv.db.StoreVoter(tx, voter)
		if HTTPErrorWithRollback(tx, w, http.StatusInternalServerError, err) {
			return
		}
	} else {
		log.Println("user declined: ", registrationReq.Username)
	}

	// commit the transaction
	err = tx.Commit()
	if HTTPErrorWithRollback(tx, w, http.StatusInternalServerError, err) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}
