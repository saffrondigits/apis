package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:abcd@localhost:1234/goclass?sslmode=disable")
	if err != nil {
		log.Println("Cannot connect to the db")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Cannot ping to the db")
		return nil, err
	}

	return db, nil
}

func main() {
	r := gin.Default()

	r.GET("/ping", ping)
	r.POST("/register", registerUser)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		log.Println("cannot start the server!")
	}
}

func registerUser(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "json data is not correct"})
		return
	}

	// Connect to the database
	_, err = connectToDatabase()
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	// Check if the email already exist in the db

	// Check if the username already exist in the db

	// Store the data into the database
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
