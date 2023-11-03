package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/saffrondigits/apis/auth"
	"github.com/saffrondigits/apis/modules"
)

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
	r.POST("/login", loginUser)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		log.Println("cannot start the server!")
	}
}

func loginUser(c *gin.Context) {
	var loginUser modules.LoginUser

	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(400, gin.H{"error": "json data is not correct"})
		return
	}

	// Connect to the database
	sql, err := connectToDatabase()
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	// Check if the username already exist in the db
	dbUser := modules.LoginUser{}
	err = sql.QueryRow("SELECT username,password FROM users where username=$1", loginUser.UserName).Scan(&dbUser.UserName, &dbUser.Password)
	if err != nil {
		log.Printf("username doesn't exist")
		c.JSON(400, gin.H{"error": "username doesn't  exist"})
		return
	}

	// Check if password matches
	if dbUser.Password != loginUser.Password {
		log.Printf("username password doesn't match")
		c.JSON(400, gin.H{"error": "username password doesn't match"})
		return
	}

	// Generate a token and return
	token, err := auth.CreateToken(loginUser.UserName)
	if err != nil {
		log.Printf("cannot create a token")
		c.JSON(500, gin.H{"error": "cannot create a token"})
		return
	}

	// Return a session data
	c.JSON(200, gin.H{"token": token})
}

func registerUser(c *gin.Context) {
	var user modules.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "json data is not correct"})
		return
	}

	// Connect to the database
	sql, err := connectToDatabase()
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	// Check if the email already exist in the db
	dbUser := modules.User{}
	err = sql.QueryRow("SELECT first_name,last_name FROM users where email=$1", user.Email).Scan(&dbUser.FirstName, &dbUser.LastName)
	if err == nil {
		log.Printf("email already exist")
		c.JSON(500, gin.H{"error": "email already exist"})
		return
	}

	// Check if the username already exist in the db
	err = sql.QueryRow("SELECT first_name,last_name FROM users where username=$1", user.UserName).Scan(&dbUser.FirstName, &dbUser.LastName)
	if err == nil {
		log.Printf("username already exist")
		c.JSON(500, gin.H{"error": "username already exist"})
		return
	}

	// Store the data into the database
	_, err = sql.Query("INSERT INTO users (first_name,last_name, email, username, password) VALUES($1,$2,$3,$4,$5)", user.FirstName, user.LastName, user.Email, user.UserName, user.Password)
	if err != nil {
		log.Printf("cannot insert the value to the database: %+v", err)
		c.JSON(500, gin.H{"error": "failed to register"})
		return
	}

	c.JSON(200, gin.H{"error": "successfully registered"})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
