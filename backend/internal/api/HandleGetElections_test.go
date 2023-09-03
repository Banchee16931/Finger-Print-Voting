package api_test

import (
	"encoding/json"
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/test/testutils"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that Server.HandleGetElections will correctly return elections gotten from the database
func TestHandleGetElections(t *testing.T) {
	t.Parallel()
	// setting up database expectations
	dbClient := testutils.MockDB{}

	dbClient.On("GetElections").Return([]types.Election{
		{
			ElectionID: 0,
			Start:      "2023-08-01",
			End:        "2023-08-02",
			Location:   "Location",
		}, {
			ElectionID: 1,
			Start:      "2023-06-01",
			End:        "2023-06-02",
			Location:   "Other Location",
		},
	}, nil)
	dbClient.On("GetResults", 0).Return([]types.Result{
		{
			ResultID:    0,
			ElectionID:  0,
			FirstName:   "CandidateOne",
			LastName:    "CandidateOne Last",
			Party:       "Example Party",
			PartyColour: "#ff000",
			Votes:       42,
		}, {
			ResultID:    1,
			ElectionID:  0,
			FirstName:   "CandidateTwo",
			LastName:    "CandidateTwo Last",
			Party:       "Other Party",
			PartyColour: "#ffff0",
			Votes:       23,
		},
	}, nil)
	dbClient.On("GetResults", 1).Return([]types.Result{}, nil)
	dbClient.On("GetCandidates", 1).Return([]types.Candidate{
		{
			CandidateID: 0,
			ElectionID:  1,
			FirstName:   "CandidateThree",
			LastName:    "CandidateThree Last",
			Party:       "Third Party",
			PartyColour: "#0000ff",
			Photo:       "",
		},
		{
			CandidateID: 1,
			ElectionID:  1,
			FirstName:   "CandidateFour",
			LastName:    "CandidateFour Last",
			Party:       "Final Party",
			PartyColour: "#00ffff",
			Photo:       "",
		},
	}, nil)
	dbClient.On("GetVotes", 1).Return([]types.Vote{
		{
			Username:    "User1",
			ElectionID:  1,
			CandidateID: 1,
		}, {
			Username:    "User2",
			ElectionID:  1,
			CandidateID: 0,
		}, {
			Username:    "User3",
			ElectionID:  1,
			CandidateID: 0,
		}, {
			Username:    "User4",
			ElectionID:  1,
			CandidateID: 0,
		}, {
			Username:    "User5",
			ElectionID:  1,
			CandidateID: 1,
		}, {
			Username:    "User6",
			ElectionID:  1,
			CandidateID: 0,
		},
	}, nil)

	// expected output
	outputState := []types.ElectionState{
		{
			ElectionID: 0,
			Start:      "2023-08-01",
			End:        "2023-08-02",
			Location:   "Location",
			Result: []types.CandidateVotes{
				{
					FirstName:   "CandidateThree",
					LastName:    "CandidateThree Last",
					Party:       "Third Party",
					PartyColour: "#0000ff",
					Votes:       42,
				}, {
					FirstName:   "CandidateTwo",
					LastName:    "CandidateTwo Last",
					Party:       "Other Party",
					PartyColour: "#ffff0",
					Votes:       23,
				},
			},
		}, {
			ElectionID: 1,
			Start:      "2023-06-01",
			End:        "2023-06-02",
			Location:   "Other Location",
			Result: []types.CandidateVotes{
				{
					FirstName:   "CandidateOne",
					LastName:    "CandidateOne Last",
					Party:       "Example Party",
					PartyColour: "#ff000",
					Votes:       4,
				}, {
					FirstName:   "CandidateFour",
					LastName:    "CandidateFour Last",
					Party:       "Final Party",
					PartyColour: "#00ffff",
					Votes:       2,
				},
			},
		},
	}

	// setting up server
	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Act
	srv.HandleGetElections(w, r)

	// Assert
	res := w.Result()

	dbClient.AssertExpectations(t)

	// check response status code
	if !assert.Equal(t, http.StatusOK, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}

	responseBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.NoError(t, err, "reading body caused an error")
	t.Logf("Response Body: %s", string(responseBody))

	var actualOupoutElectionState []types.ElectionState
	assert.NoError(t, json.Unmarshal(responseBody, &actualOupoutElectionState), "failed to read response body")

	if assert.Equal(t, len(actualOupoutElectionState), len(outputState), "len of actual did not match expected") {
		return
	}

	sort.SliceStable(outputState, func(i, j int) bool {
		return outputState[i].ElectionID < outputState[j].ElectionID
	})

	sort.SliceStable(actualOupoutElectionState, func(i, j int) bool {
		return actualOupoutElectionState[i].ElectionID < actualOupoutElectionState[j].ElectionID
	})

	for i := 0; i < len(actualOupoutElectionState); i++ {
		assert.Equal(t, actualOupoutElectionState[i].ElectionID, outputState[i].ElectionID, "election IDs don't match")
		assert.Equal(t, actualOupoutElectionState[i].Start, outputState[i].Start, "starts don't match")
		assert.Equal(t, actualOupoutElectionState[i].End, outputState[i].End, "ends don't match")
		assert.Equal(t, actualOupoutElectionState[i].Location, outputState[i].Location, "location don't match")
		assert.ElementsMatch(t, actualOupoutElectionState[i].Result, outputState[i].Result, "results don't match")
	}
}

// Tests that Server.HandleGetElections will correctly return elections gotten from the database, when there is no elections in the database
func TestHandleGetElections_NoElections(t *testing.T) {
	t.Parallel()
	registrationsInDB := []types.Registrant{}

	// setting up database expectations
	dbClient := testutils.MockDB{}

	dbClient.On("GetRegistrants").Return(registrationsInDB, nil)

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Act
	srv.HandleGetRegistrations(w, r)

	// Assert
	res := w.Result()

	dbClient.AssertExpectations(t)

	// check response status code
	if !assert.Equal(t, http.StatusOK, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}

	var body []byte
	_, err := res.Body.Read(body)
	assert.NoError(t, err, "failed to read response body")

	defer res.Body.Close()

	assert.Equal(t, "", string(body), "response body was not what was expeceted")
}

// Tests that if the database runs into an error it is correctly reported by HandleGetElections
func TestHandleGetElections_DBError(t *testing.T) {
	t.Parallel()

	// setting up database expectations
	dbClient := testutils.MockDB{}

	dbClient.On("GetRegistrants").Return([]types.Registrant{}, fmt.Errorf("I am an error"))

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Act
	srv.HandleGetRegistrations(w, r)

	// Assert
	dbClient.AssertExpectations(t)

	res := w.Result()
	if !assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "incorrect status code") {
		return
	}

	AssertHTTPErrorResponse(t, res.Body)
}
