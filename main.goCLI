package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	// Define the path where the DuckDB file will be saved
	dbPath := "mydatabase.duckdb"

	// Check if the file already exists and remove it for this example
	if _, err := os.Stat(dbPath); err == nil {
		if err := os.Remove(dbPath); err != nil {
			log.Fatalf("Failed to remove existing database file: %v", err)
		}
		fmt.Println("Removed existing database file")
	}

	// Connect to DuckDB (this will create the file if it doesn't exist)
	db, err := sql.Open("duckdb", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to DuckDB database")

	// Create a table
	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Created 'users' table")

	// Insert some sample data
	_, err = db.Exec(`
		INSERT INTO users (id, name, email) VALUES 
		(1, 'John Doe', 'john@example.com'),
		(2, 'Jane Smith', 'jane@example.com'),
		(3, 'Bob Johnson', 'bob@example.com')
	`)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
	fmt.Println("Inserted sample data")

	// Query the data to verify
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}
	defer rows.Close()

	fmt.Println("\nUsers in the database:")
	fmt.Println("----------------------")
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}

	fmt.Printf("\nDuckDB database saved to: %s\n", dbPath)
}
