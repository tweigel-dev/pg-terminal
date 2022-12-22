package test
func TestGetSQLStatement(t *testing.T) {
	// Set up mock command object
	cmd := &cobra.Command{}
	cmd.Flags().String("file", "test.sql", "")
	cmd.Flags().String("env", "", "")

	// Test getting SQL statement from file
	sql, err := getSQLStatement(cmd)
	if err != nil {
		t.Errorf("Error getting SQL statement from file: %v", err)
	}
	if sql != "SELECT * FROM test_table;" {
		t.Errorf("Unexpected SQL statement: %s", sql)
	}

	// Set up mock command object with environment variable
	cmd = &cobra.Command{}
	cmd.Flags().String("file", "", "")
	cmd.Flags().String("env", "TEST_SQL", "")
	os.Setenv("TEST_SQL", "SELECT * FROM test_table;")

	// Test getting SQL statement from environment variable
	sql, err = getSQL Statement(cmd)
	if err != nil {
		t.Errorf("Error getting SQL statement from environment variable: %v", err)
	}
	if sql != "SELECT * FROM test_table;" {
		t.Errorf("Unexpected SQL statement: %s", sql)
	}
	os.Unsetenv("TEST_SQL")

	// Set up mock command object with no file or environment variable
	cmd = &cobra.Command{}
	cmd.Flags().String("file", "", "")
	cmd.Flags().String("env", "", "")

	// Test getting SQL statement with no file or environment variable
	_, err = getSQLStatement(cmd)
	if err == nil {
		t.Error("Expected error getting SQL statement with no file or environment variable")
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