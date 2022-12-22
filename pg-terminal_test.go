package main
import (
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"testing"
)

func TestGetSQLStatementPath(t *testing.T) {
	// Set up mock command object
	cmd := &cobra.Command{}
	cmd.Flags().String("file", "example.sql", "")
	// Test getting SQL statement from file
	sql, err := getSQLStatement(cmd)
	if err != nil {
		t.Errorf("Error getting SQL statement from file: %v", err)
	}
	if sql == "" {
		t.Errorf("no sql statement extracted from file %s", sql)
	}

	
}
func TestGetSQLStatementEnv(t *testing.T) {
	// Set up mock command object
	os.Setenv("PG_TERMINAL_SQL", "CREATE or REPLACE TABLE lesson (student VARCHAR(255),room INT);")
	cmd := &cobra.Command{}
	// Test getting SQL statement from file
	sql, err := getSQLStatement(cmd)
	if err != nil {
		t.Errorf("Error getting SQL statement from file: %v", err)
	}
	if sql == "" {
		t.Errorf("Unexpected SQL statement: %s", sql)
	}
	
}

func TestGetConnStr(t *testing.T) {
	// Set up mock command object
	cmd := &cobra.Command{}
	cmd.Flags().String("user", "test_user", "")
	cmd.Flags().String("database", "test_database", "")
	cmd.Flags().String("password", "test_password", "")
	cmd.Flags().Int("port", 5432, "")
	cmd.Flags().String("ip", "127.0.0.1", "")

	// Test building connection string
	connStr, err := getConnStr(cmd)
	if err != nil {
		t.Errorf("Error building connection string: %v", err)
	}
	expected := "user=test_user password=test_password dbname=test_database port=5432 host=127.0.0.1 sslmode=disable"
	if connStr != expected {
		t.Errorf("Unexpected connection string: %s", connStr)
	}
}