package database

import (
	"InsultAPI/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
)

const EnvPath = "DATABASE_PATH"
const TableName = "insults"

var db *sql.DB

func InitDB() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, using environment variables instead.")
	}

	dbPath := viper.GetString("database_path")
	if dbPath == "" {
		dbPath = "insultAPI.db"
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Ensure the table exists
	if !tableExists(db) {
		err = createTable()
		if err != nil {
			log.Fatal("Failed to create table:", err)
		}
		fmt.Println("Table created")
	}
}

func tableExists(db *sql.DB) bool {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?;"
	var name string
	err := db.QueryRow(query, TableName).Scan(&name)
	return err == nil
}

func createTable() error {
	query := "CREATE TABLE insults (id INTEGER PRIMARY KEY AUTOINCREMENT, insult TEXT NOT NULL)"
	_, err := db.Exec(query)
	return err
}

func GetRandomInsult() (models.Insult, error) {
	var insult models.Insult
	query := "SELECT id, insult FROM insults ORDER BY RANDOM() LIMIT 1;"
	err := db.QueryRow(query).Scan(&insult.ID, &insult.Insult)
	if err != nil {
		return models.Insult{}, fmt.Errorf("failed to fetch insult: %w", err)
	}
	return insult, nil
}

func GetAllInsults() (models.Insults, error) {
	rows, err := db.Query("SELECT id, insult FROM insults")
	if err != nil {
		return nil, fmt.Errorf("failed to obtain all insults: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var insults models.Insults
	for rows.Next() {
		var insult models.Insult
		if err := rows.Scan(&insult.ID, &insult.Insult); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		insults = append(insults, insult)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return insults, nil
}

func AddInsult(text string) error {
	_, err := db.Exec("INSERT INTO insults (insult) VALUES (?)", text)
	if err != nil {
		return err
	}
	return nil
}

func DeleteInsult(insultID int) error {
	// Check if the insult exists first
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM insults WHERE id = ?)", insultID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("insult not found")
	}

	_, err = db.Exec("DELETE FROM insults WHERE id = ?", insultID)
	if err != nil {
		return fmt.Errorf("failed to delete insult: %w", err)
	}
	return nil
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			return
		}
	}
}
