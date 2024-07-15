package controllers

import (
  "fmt"
  "time"

  "github.com/TravisLinkey/knowledge-notification-system/models"
)

func UpdateNote(id int64, location string, reviewed int) error {
  note, err := models.GetNoteById(id)
  if err != nil {
    fmt.Println("Failed to get the note id=%s, %v", id, err)
    return err
  }

  note.Location = location
  note.Reviewed = reviewed

  err = note.Update()
  if err != nil {
    fmt.Println("Failed to update note id=%s, %v", id, err)
    return err
  }

  fmt.Println("Note updated successfully!")
  return nil
}

func CreateNote(id int64, location string) {
  note := models.Note{
     ID: id,
     Created: time.Now(),
     Reviewed: 0,
     Location: location,
   } 

   err := note.Save()
   if err != nil {
     fmt.Println("Some error, %v", err)
    return
   }
   fmt.Print("Saved the note!")
}

func FetchAllNotes() {
  allNotes, _ := models.GetAllNotes()

  fmt.Println("All Notes: ")
  for i, s := range allNotes {
    fmt.Println(i, s)
  }
}

func GetNote(id int64) (*models.Note, error) {
  singleNote, err := models.GetNoteById(id)
  if err != nil {
    fmt.Println("Failed to retrieve single note: %v", err)
    return nil, err
  }

  return singleNote, nil
}

