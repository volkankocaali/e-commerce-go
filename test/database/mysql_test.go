package database_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	"log"
	"testing"
)

func TestMySQLConnection(t *testing.T) {
	cfg := config.Config{
		MySQLHost:     "localhost",
		MySQLPort:     "3304",
		MySQLUser:     "root",
		MySQLPassword: "secret",
		MySQLDatabase: "e_commerce",
	}

	db, err := database.NewMysqlDB(cfg)

	if err != nil {
		t.Fatalf("Failed to create MySQL connection: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to close MySQL connection: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.NoError(t, err, "Failed to create MySQL connection")
}
