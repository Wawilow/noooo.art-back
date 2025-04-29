package db

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// Optional: For logging SQL
)

type Articles struct {
	ID          string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Match UUID type
	CreatedAt   time.Time  `json:"created_at"`
	Header      *string    `json:"header"`
	HeaderLink  *string    `json:"header_link"`
	SubHeader   *string    `json:"sub_header"`
	Author      *string    `json:"author"`
	Date        *time.Time `json:"date"`
	Description *string    `json:"description"`
}

func (u *Articles) TableName() string {
	// custom table name, this is default
	return "prod.articles"
}

type Grm struct {
	DB *gorm.DB
}

func NewGrm(dsn string) (*Grm, error) {
	if dsn == "" {
		log.Fatal("Error: SUPABASE_DB_DSN environment variable must be set.")
		return nil, errors.New("")
	}

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("Successfully connected to Supabase database using GORM!")
	return &Grm{DB: db}, nil
}

func (g *Grm) GetArticles() (*[]Articles, error) {
	var foundArticle []Articles

	r := g.DB.Raw("SELECT * FROM prod.articles").Scan(&foundArticle)
	if r.Error != nil {
		return nil, r.Error
	}

	return &foundArticle, nil
}
