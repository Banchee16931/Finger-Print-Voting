package database

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
)

func (client *Client) StoreRegistrant(registrant types.Registrant) error {
	return cerr.ErrUnimplemented
}

func (client *Client) GetRegistrants() ([]types.Registrant, error) {
	return []types.Registrant{}, cerr.ErrUnimplemented
}

func (client *Client) DeleteRegistrant(registrantID int) error {
	return cerr.ErrUnimplemented
}
