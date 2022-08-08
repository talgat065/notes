package main

import (
	"fmt"
	"log"
	"notes/internal/application/services"
	"notes/internal/application/services/dto"
	"notes/internal/infrastructure/repositories"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/google/uuid"
)

func main() {
	repository := repositories.NewNotesMemoryRepository()
	service := services.NewNotesService(repository)

	app := &cli.App{
		Name:  "note",
		Usage: "note",
		Action: func(*cli.Context) error {
			fmt.Println("creating note...")
			request := dto.CreateNoteRequest{ID: uuid.NewString(), Text: "#music The Libertines Death On Stairs"}
			id, err := service.CreateNote(request)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				fmt.Printf("note created: %s\n", id)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
