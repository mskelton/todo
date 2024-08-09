package storage

import (
	"errors"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func getDBPath() (string, error) {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("invalid database: unable to find home directory")
	}

	dir := filepath.Join(home, ".local", "state", "todo")

	// Auto-create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}

	return filepath.Join(dir, "todo.db"), nil
}

var db *gorm.DB

func GetDB() (*gorm.DB, error) {
	// Return the cached DB if available
	if db != nil {
		return db, nil
	}

	// Get the DB path
	path, err := getDBPath()
	if err != nil {
		return nil, errors.New("Invalid database path")
	}

	// Open a new connection to the DB
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	return db, err
}

func Migrate() error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	return errors.Join(
		db.AutoMigrate(&Project{}),
		db.AutoMigrate(&Task{}),
	)
}
