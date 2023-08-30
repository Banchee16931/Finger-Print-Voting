package database

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
)

func (client *Client) StoreElection(election types.Election) error {
	return cerr.ErrUnimplemented
}

func (client *Client) GetElections() ([]types.Election, error) {
	return []types.Election{}, cerr.ErrUnimplemented
}

func (client *Client) DeleteCandidates(electionID int) error {
	return cerr.ErrUnimplemented
}
