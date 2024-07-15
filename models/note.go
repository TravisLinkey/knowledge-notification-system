package models

import (
  "time"
  "fmt"

  "github.com/TravisLinkey/knowledge-notification-system/db"
)

type Note struct {
  ID        int64         `json:"id"`
  Created   time.Time   `json:"created"`
  Reviewed  int        `json:"reviewed"`
  Location  string      `json:"location"`
}

var notes = []Note{}

func (note Note) Save() error {
  query := `
    INSERT INTO notes(id, created, reviewed, location) 
    VALUES (?,?,?,?)
  `

  statement, err := db.Database.Prepare(query)
  if err != nil {
    fmt.Println("Failed to create INSERT statement, %v", err)
    return err
  }
  
  defer statement.Close()
  result, err := statement.Exec(note.ID, note.Created, note.Reviewed, note.Location)
  if err != nil {
    fmt.Println("Failed to execute insert statment, %v", err)
    return err
  }

  id, err := result.LastInsertId()
  note.ID = id
  return nil
}

func (note Note) Update() error {
  query := `
    UPDATE notes
    SET reviewed = ?, location = ?
    WHERE id = ?
  `

  statement, err := db.Database.Prepare(query)
  if err != nil {
    fmt.Println("Failed to create UPDATE statement, %v", err)
    return err
  }

  defer statement.Close()

  _, err = statement.Exec(note.Reviewed, note.Location, note.ID)
  return err
}

func GetAllNotes() ([]Note, error) {
  query := `SELECT * FROM notes`
  rows, err := db.Database.Query(query)
  if err != nil {
    fmt.Println("Failed to retrieve rows, %v", err)
    return nil, err
  }
  defer rows.Close()

  var notes []Note

  for rows.Next() {
    var note Note
      err := rows.Scan(&note.ID, &note.Created, &note.Reviewed, &note.Location)

      if err != nil {
        fmt.Println("Failed to scan row: %v", err)
        return nil, err
      }

      notes = append(notes, note)
  }

  return notes, nil
}

func GetNoteById(id int64) (*Note, error) {
  query := `
    SELECT * FROM notes
    WHERE id = ?
  `

  row := db.Database.QueryRow(query, id)

  var note Note
  err := row.Scan(&note.ID, &note.Created, &note.Reviewed, &note.Location)
  if err != nil {
    return nil, err
  }

  return &note, nil
}

func GetUnreviewedNotes() ([]Note, error) {
  query := `
    SELECT * FROM notes
    WHERE reviewed = 0
  `

  rows, err := db.Database.Query(query)
  if err != nil {
    fmt.Println("Failed to retrieve rows, %v", err)
    return nil, err
  }
  defer rows.Close()

  var notes []Note

  for rows.Next() {
    var note Note
      err := rows.Scan(&note.ID, &note.Created, &note.Reviewed, &note.Location)

      if err != nil {
        fmt.Println("Failed to scan row: %v", err)
        return nil, err
      }

      notes = append(notes, note)
  }

  return notes, nil
}
