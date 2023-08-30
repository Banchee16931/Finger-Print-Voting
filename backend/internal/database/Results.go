package database

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
)

func (client *Client) StoreResult(result types.Result) error {
	return cerr.ErrUnimplemented
}

func (client *Client) GetResults(electionID int) ([]types.Result, error) {
	return []types.Result{}, cerr.ErrUnimplemented
}
