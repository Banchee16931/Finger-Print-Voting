package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Client struct {
	db *sql.DB
}

func NewDatabase(cfg config.DBConfig) (*Client, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Server, cfg.Port, cfg.Username,
		cfg.Password, cfg.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	client := Client{
		db: db,
	}

	return &client, nil
}

func NewClientFromDatabase(db *sql.DB) *Client {
	client := Client{
		db: db,
	}

	return &client
}

func (client *Client) EnsureValidSchema(schemaLocation string) error {
	setup, err := client.IsSchemaSetup()
	if err != nil {
		return err
	}

	if !setup {
		if err := client.SetupSchema(schemaLocation); err != nil {
			return err
		}
	}

	return nil
}

func (client *Client) Close() error {
	return client.db.Close()
}

func (client *Client) Begin() (*sql.Tx, error) {
	return client.db.Begin()
}

func (client *Client) DropDBTables() error {
	_, err := client.db.Exec(`DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;`)
	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err)
	}

	log.Println("Dropped all tables")
	return nil
}

func (client *Client) Raw() *sql.DB {
	return client.db
}
