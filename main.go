package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type ToDo struct {
	ID        uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"isDone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	// Get database environment variables.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Try connecting to the database.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&ToDo{})

	db.Create(&ToDo{Content: "Study Go.", IsDone: true})

	// Initialize app router.
	router := gin.Default()
	router.GET("/todos", getToDos)

	router.Run("localhost:8080")
}

func getToDos(c *gin.Context) {
	var toDos []ToDo
	db.Find(&toDos)
	c.IndentedJSON(http.StatusOK, toDos)
}
