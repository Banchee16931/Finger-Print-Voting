package database

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
)

func (client *Client) StoreVote(vote types.Vote) error {
	return cerr.ErrUnimplemented
}

func (client *Client) GetVotes(electionID int) ([]types.Vote, error) {
	return []types.Vote{}, cerr.ErrUnimplemented
}

func (client *Client) DeleteVotes(electionID int) error {
	return cerr.ErrUnimplemented
}
