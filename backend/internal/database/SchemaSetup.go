package database

import (
	"finger-print-voting-backend/internal/cerr"
	"fmt"
	"log"
	"os"
)

var SQLTables = []string{"voter_details", "users", "registrants"}

func (client *Client) IsSchemaSetup() (bool, error) {
	schemaSetup := true

	for i := 0; i < len(SQLTables); i++ {
		log.Printf("Checking for table: %s", SQLTables[i])
		row := client.db.QueryRow(`SELECT EXISTS 
		(SELECT * FROM INFORMATION_SCHEMA.TABLES 
		WHERE table_name=$1)`, SQLTables[i])

		var tableExists bool = false

		err := row.Scan(&tableExists)
		if err != nil {
			return false, fmt.Errorf("%w: failed to scan db from query: %s", cerr.ErrDB, err.Error())
		}

		if !tableExists {
			log.Println("Table doesn't exist")
			schemaSetup = false
		}
	}

	return schemaSetup, nil
}

func (client *Client) SetupSchema(schemaLocation string) error {
	dirEntries, err := os.ReadDir(schemaLocation)
	if err != nil {
		return err
	}

	for i := 0; i < len(dirEntries); i++ {
		schemaEntry := fmt.Sprintf("%s\\%s", schemaLocation, dirEntries[i].Name())
		log.Printf("Reading Schema: %s\n", schemaEntry)
		sql, err := os.ReadFile(schemaEntry)
		if err != nil {
			return err
		}

		_, err = client.db.Exec(string(sql))
		if err != nil {
			return fmt.Errorf("%w: %s", cerr.ErrDB, err)
		}
	}

	return nil
}
