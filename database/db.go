package database

import (
	"InsultAPI/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var tableName = "insults"
var db *sql.DB

func InitDB() {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath != "" {
		dbPath = dbPath + "/insultAPI.db"
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Ensure the table exists
	if !tableExists(db, tableName) {
		err = createTable()
		if err != nil {
			log.Fatal("Failed to create table:", err)
		}
		fmt.Println("Table created")
	}
}

func tableExists(db *sql.DB, tableName string) bool {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?;"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	return err == nil
}

func createTable() error {
	_, err := db.Exec(`
        CREATE TABLE insults (
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            insult TEXT NOT NULL
        )
    `)
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
	var insults models.Insults

	for rows.Next() {
		var insult models.Insult
		if err := rows.Scan(&insult.ID, &insult.Insult); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		insults = append(insults, insult)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
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
	return err
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			return
		}
	}
}
