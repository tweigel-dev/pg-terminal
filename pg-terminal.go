package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "v0.0.4",
	Use:   "pg-terminal",
	Short: "A tool for sending SQL statements to a PostgreSQL instance",
	Long: `pg-terminal is a command-line tool for sending SQL statements to a 
PostgreSQL instance using the pq library. It can accept a SQL statement as 
a file or environment variable and uses flags to specify the connection 
details for the PostgreSQL instance.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get SQL statement from file or environment variable
		sqlStatement, err := getSQLStatement(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get connection details from flags
		connStr, err := getConnStr(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Connect to PostgreSQL instance
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer db.Close()

		// Execute SQL statement
		_, err = db.Exec(sqlStatement)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func main() {
	// Add flags for connection details
	rootCmd.Flags().StringP("file", "f", "", "path to the sql file")
	rootCmd.Flags().StringP("user", "u", "", "Username for the PostgreSQL instance")
	rootCmd.Flags().StringP("database", "d", "", "Database name for the PostgreSQL instance")
	rootCmd.Flags().StringP("password", "p", "", "Password for the PostgreSQL instance")
	rootCmd.Flags().IntP("port", "P", 5432, "Port for the PostgreSQL instance")
	rootCmd.Flags().StringP("ip", "i", "localhost", "IP address for the PostgreSQL instance")
	rootCmd.MarkFlagRequired("user")
	rootCmd.MarkFlagRequired("database")
	rootCmd.MarkFlagRequired("password")

	// Parse flags and execute command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getSQLStatement(cmd *cobra.Command) (string, error) {
	// Check if SQL statement is provided as file
	if sqlFile, err := cmd.Flags().GetString("file"); sqlFile != "" && err == nil {
		// Read SQL statement from file
		bytes, err := ioutil.ReadFile(sqlFile)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	// Check if SQL statement is provided as environment variable
	if sqlStatement, ok := os.LookupEnv("PG_TERMINAL_SQL"); ok {
		// Get SQL statement from environment variable
		return sqlStatement, nil
	} else {
		return "", fmt.Errorf("Environment variable '%s' not set", sqlStatement)
	}

	// Return error if SQL statement not provided as file or environment variable
	return "", fmt.Errorf("No SQL statement provided")
}

func getConnStr(cmd *cobra.Command) (string, error) {
	// Get connection details from flags
	user, err := cmd.Flags().GetString("user")
	if err != nil {
		return "", err
	}
	database, err := cmd.Flags().GetString("database")
	if err != nil {
		return "", err
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return "", err
	}
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return "", err
	}
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return "", err
	}

	// Build connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%d host=%s sslmode=disable",
		user, password, database, port, ip)
	return connStr, nil
}
