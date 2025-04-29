package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"noooo.art/app/config"
	"noooo.art/app/db"
)

func handleStatus(c *fiber.Ctx) error {
	response := fiber.Map{
		"status": "success",
	}

	return c.JSON(response)
}

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	db, err := db.NewGrm(*config.SUPABASE_DSN)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	app := fiber.New()

	app.Get("/api/status", handleStatus)
	app.Get("/api/content", func(c *fiber.Ctx) error {
		content, err := db.GetArticles()
		if err != nil {
			return err
		}

		type Cont struct {
			Header      *string    `json:"header,omitempty"`
			HeaderLink  *string    `json:"header_link,omitempty"`
			SubHeader   *string    `json:"sub_header,omitempty"`
			Author      *string    `json:"author,omitempty"`
			Date        *time.Time `json:"date,omitempty"`
			Description *string    `json:"description,omitempty"`
		}
		r := []Cont{}

		for _, i := range *content {
			r = append(r, Cont{
				Header:      i.Header,
				HeaderLink:  i.HeaderLink,
				SubHeader:   i.SubHeader,
				Author:      i.Author,
				Date:        i.Date,
				Description: i.Description,
			})
		}

		return c.JSON(r)
	})

	log.Println("Starting server on :3000...")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
