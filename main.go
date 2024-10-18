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
	var err error

	// Get database environment variables.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Try connecting to the database.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema.
	db.AutoMigrate(&ToDo{})

	// Populate sample data.
	// db.Create(&ToDo{Content: "Study Go.", IsDone: true})

	// Initialize app router.
	router := gin.Default()
	router.GET("/todos", getToDos)
	router.POST("/todo", createToDo)
	router.GET("/todo/:id", getToDo)
	router.PUT("/todo/:id", updateToDo)
	router.DELETE("/todo/:id", deleteToDo)

	router.Run("localhost:8080")
}

func getToDos(c *gin.Context) {
	var toDos []ToDo
	db.Find(&toDos)
	c.IndentedJSON(http.StatusOK, toDos)
}

func createToDo(c *gin.Context) {
	var toDo ToDo
	c.BindJSON(&toDo)
	db.Create(&toDo)
	c.IndentedJSON(http.StatusOK, toDo)
}

func getToDo(c *gin.Context) {
	// Theses parameters are path parameters.
	id := c.Param("id")
	var toDo ToDo

	// Without column name specified, GORM uses primary key which is ID.
	db.First(&toDo, id)
	c.IndentedJSON(http.StatusOK, toDo)
}

func updateToDo(c *gin.Context) {
	id := c.Param("id")
	var toDo ToDo
	db.First(&toDo, id)
	c.BindJSON(&toDo)
	db.Save(&toDo)
	c.IndentedJSON(http.StatusOK, toDo)
}

func deleteToDo(c *gin.Context) {
	id := c.Param("id")
	var toDo ToDo
	db.First(&toDo, id)
	db.Delete(&toDo)
	c.IndentedJSON(http.StatusOK, toDo)
}
