package main

import (
  "fmt"

  "github.com/TravisLinkey/knowledge-notification-system/controllers"
  "github.com/TravisLinkey/knowledge-notification-system/models"
)


func main() {
  fmt.Println("-- Application starting --")

  // GET SINGLE NOTE
  // singleNote, _ := controllers.GetNote(3)
  // fmt.Println("Single Note: ", singleNote)

  // UPDATE SINGLE NOTE

  // CREATE NOTE
  controllers.CreateNote(4, "Location_4")
  controllers.CreateNote(5, "Location_5")
  controllers.CreateNote(6, "Location_6")
  controllers.CreateNote(7, "Location_7")

  _ = controllers.UpdateNote(6, "NewLocation", 1)

  unreviewed, _ := models.GetUnreviewedNotes()
  for i, s := range unreviewed {
    fmt.Println(i, s)
  }
}
