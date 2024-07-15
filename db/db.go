package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "fmt"
)

var Database *sql.DB

func init() {
  InitDB()
  CreateTables()
}

func InitDB() {
  var err error

  Database, err = sql.Open("sqlite3", "./api.db")
  if err != nil {
    fmt.Println("Error: ", err)
    panic("Could not connect to database.")
  }

  Database.SetMaxOpenConns(10)
  Database.SetMaxIdleConns(5)

  
    // Ping to verify connection is established
    if err = Database.Ping(); err != nil {
      fmt.Println("failed to verify connection: %w", err)
      panic("failed to verify connection")
    }
}

func CreateTables() {
  if Database == nil {
    fmt.Println("Database not initialized")
    panic("failed to verify connection")
  }

  createTable := `
    CREATE TABLE IF NOT EXISTS notes (
      id INTEGER PRIMARY KEY,
      created DATETIME,
      reviewed INTEGER,
      location TEXT
  )`

  _, err := Database.Exec(createTable)
  if err != nil {
    fmt.Printf("failed to create table %w", err)
    panic("failed to create table")
  }

  fmt.Println("-- Table created --")
}

