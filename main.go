package main

import (
	"log"
	"time"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"noooo.art/app/config"
	"noooo.art/app/db"
)

func handleStatus(c *fiber.Ctx) error {
	response := fiber.Map{
		"status": "success",
		"version": "v0.0.3",
	}

	return c.JSON(response)
}

func stringPtr(s string) *string {
	return &s
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
	fmt.Println(db)

	app.Get("/api/status", handleStatus)
	app.Get("/api/content", func(c *fiber.Ctx) error {
		// content, err := db.GetArticles()
		// if err != nil {
		// 	return err
		// }

		type Cont struct {
			Header      *string    `json:"header,omitempty"`
			HeaderLink  *string    `json:"header_link,omitempty"`
			SubHeader   *string    `json:"sub_header,omitempty"`
			Author      *string    `json:"author,omitempty"`
			Date        *time.Time `json:"date,omitempty"`
			Description *string    `json:"description,omitempty"`
		}
		r := []Cont{}

		// for _, i := range *content {
		// 	r = append(r, Cont{
		// 		Header:      i.Header,
		// 		HeaderLink:  i.HeaderLink,
		// 		SubHeader:   i.SubHeader,
		// 		Author:      i.Author,
		// 		Date:        i.Date,
		// 		Description: i.Description,
		// 	})
		// }
		
		now := time.Now() // Get current time and store in a variable
		descText := `Somewhere deep, deep down in a k-hole, Yung Lean and Sam Bankman-Fried are holding hands, thinking of stuff like this. Which is to say, while Rafman’s show is symptomatic of “internet culture,” and duly hypnotic, it doesn’t transcend the guise of innovation. Now, do I like the smoking angel-deer? Of course. But the show never answers Why. Why, apart from “AI,” should we care? What’s in it for us to invest in these people, that sequencing, those anecdotes? When pressed, the gallery throws around ideas about worldbuilding and character. They cite Cloudy Heart, the AI popstar who “has real fans online,” implying that Rafman’s world bleeds into “the real” in prescient ways. Perhaps somewhere, but not here. These characters are silos in a vacuum – artifacts of the means of production – who lack a world built for the moments needed to bring them to life.`
		
		r = append(r, Cont{
			Header:      stringPtr("Proof of Concept"), // Get *string
			HeaderLink:  stringPtr("https://spruethmagers.com/exhibitions/jon-rafman-proof-of-concept-los-angeles/"), // Get *string
			SubHeader:   stringPtr(" – ONE POINT FIVE STARS"), // Get *string
			Author:      stringPtr("Jon Rafman"),          // Get *string
			Date:        &now,                           // Get *time.Time from variable
			Description: stringPtr(descText),              // Get *string from variable holding the multi-line text
		})

		return c.JSON(r)
	})

	log.Println("Starting server on :3000...")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
