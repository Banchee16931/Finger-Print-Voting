package database_test

import (
	"finger-print-voting-backend/internal/database"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Tests that the EnsureSchema function correctly reads the current setup of the database and then creates the schema ontop of it
func TestEnsureSchema(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}
	defer db.Close()

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	for i := 0; i < len(database.SQLTables); i++ {
		mock.ExpectQuery(`SELECT EXISTS 
		\(SELECT \* FROM INFORMATION_SCHEMA\.TABLES 
		WHERE table_name=\$1\)`).
			WithArgs(database.SQLTables[i]).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	}

	amountOfSchemaFiles := 7
	for i := 0; i < amountOfSchemaFiles; i++ {
		mock.ExpectExec(".*").WillReturnResult(sqlmock.NewResult(0, 1))
	}

	codebaseLoc, err := os.Getwd()
	assert.NoError(t, err, "failed to get working directory")

	dir := fmt.Sprintf("%s\\schemas", codebaseLoc)

	assert.NoError(t, client.EnsureValidSchema(dir), "ensure valid schema returned an error")

	assert.NoError(t, mock.ExpectationsWereMet(), "database expectations were not met")
}

// Tests that the EnsureSchema function correctly reads the current setup of the database as already setup and then doesn't attempt to set it up again
func TestEnsureSchema_SchemaInitialised(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}
	defer db.Close()

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	for i := 0; i < len(database.SQLTables); i++ {
		mock.ExpectQuery(`SELECT EXISTS 
		\(SELECT \* FROM INFORMATION_SCHEMA\.TABLES 
		WHERE table_name=\$1\)`).
			WithArgs(database.SQLTables[i]).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
	}

	codebaseLoc, err := os.Getwd()
	assert.NoError(t, err, "failed to get working directory")

	dir := fmt.Sprintf("%s\\schemas", codebaseLoc)

	assert.NoError(t, client.EnsureValidSchema(dir), "ensure valid schema returned an error")

	assert.NoError(t, mock.ExpectationsWereMet(), "database expectations were not met")
}

// Tests that an error is reported if there is a failure to read the daatbase table schema
func TestEnsureSchema_GetSchemaStateDBError(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}
	defer db.Close()

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	mock.ExpectQuery(`SELECT EXISTS 
		\(SELECT \* FROM INFORMATION_SCHEMA\.TABLES 
		WHERE table_name=\$1\)`).
		WithArgs(database.SQLTables[0]).WillReturnError(fmt.Errorf("error"))

	codebaseLoc, err := os.Getwd()
	assert.NoError(t, err, "failed to get working directory")

	dir := fmt.Sprintf("%s\\schemas", codebaseLoc)

	assert.Error(t, client.EnsureValidSchema(dir), "ensure valid schema returned an error")

	assert.NoError(t, mock.ExpectationsWereMet(), "database expectations were not met")
}

// Tests that an error is reported if there is a failure to update the daatbase table schema
func TestEnsureSchema_UpdateSchemaDBError(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}
	defer db.Close()

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	for i := 0; i < len(database.SQLTables); i++ {
		mock.ExpectQuery(`SELECT EXISTS 
		\(SELECT \* FROM INFORMATION_SCHEMA\.TABLES 
		WHERE table_name=\$1\)`).
			WithArgs(database.SQLTables[i]).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	}

	mock.ExpectExec(".*").WillReturnError(fmt.Errorf("error"))

	codebaseLoc, err := os.Getwd()
	assert.NoError(t, err, "failed to get working directory")

	dir := fmt.Sprintf("%s\\schemas", codebaseLoc)

	assert.Error(t, client.EnsureValidSchema(dir), "ensure valid schema returned an error")

	assert.NoError(t, mock.ExpectationsWereMet(), "database expectations were not met")
}

// Tests that an error is reported if the schema location given doesn't exist
func TestEnsureSchemaInvalidSchemaDirectoryError(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}
	defer db.Close()

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	for i := 0; i < len(database.SQLTables); i++ {
		mock.ExpectQuery(`SELECT EXISTS 
		\(SELECT \* FROM INFORMATION_SCHEMA\.TABLES 
		WHERE table_name=\$1\)`).
			WithArgs(database.SQLTables[i]).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	}

	codebaseLoc, err := os.Getwd()
	assert.NoError(t, err, "failed to get working directory")

	dir := fmt.Sprintf("%s\\not-schema-location", codebaseLoc)

	assert.Error(t, client.EnsureValidSchema(dir), "ensure valid schema returned an error")

	assert.NoError(t, mock.ExpectationsWereMet(), "database expectations were not met")
}

// Tests that the database is closed properally when client.Close() is ran
func TestClose(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}

	mock.ExpectClose()

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	assert.NoError(t, client.Close(), "close returned an error")
}

// Tests that an error is reported correctly if given up db.CLose()
func TestClose_Error(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}

	mock.ExpectClose().WillReturnError(fmt.Errorf("error"))

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	assert.Error(t, client.Close(), "close did not produce an error")
}

// Tests that a transacion is started properally when client.Begin() is ran
func TestBegin(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}

	mock.ExpectBegin().WillReturnError(fmt.Errorf(""))

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	_, err = client.Begin()
	assert.Error(t, err, "begin produced an error")
}

// Tests that an error is reported correctly if given up db.Begin()
func TestBegin_Error(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}

	mock.ExpectBegin().WillReturnError(fmt.Errorf(""))

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	_, err = client.Begin()
	assert.Error(t, err, "begin did not produce an error")
}

// Tests that the databases tables are dropped correctly when client.DropDBTables() is ran
func TestDropDBTable(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}

	mock.ExpectExec(`DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;`)

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	assert.Error(t, client.DropDBTables(), "DropTables produced an error")
}

// Tests that an error is reported correctly if given up db.DropDBTables()
func TestDropDBTables_DBError(t *testing.T) {
	t.Parallel()

	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %v", err)
	}

	mock.ExpectExec(`DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;`).WillReturnError(fmt.Errorf(""))

	// Create a new Client with the mock database connection
	client := database.NewClientFromDatabase(db)

	assert.Error(t, client.DropDBTables(), "DropTables did not produce an error")
}
