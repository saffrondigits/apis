package route

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saffrondigits/apis/auth"
	"github.com/saffrondigits/apis/middleware"
	"github.com/saffrondigits/apis/models"
	"github.com/sirupsen/logrus"
)

type handler struct {
	sql    *sql.DB
	logger *logrus.Logger
}

func NewHandler(db *sql.DB, log *logrus.Logger) *handler {
	return &handler{
		sql:    db,
		logger: log,
	}
}

func Route(handler *handler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.ping)

	r.POST("/register", handler.registerUser)
	r.POST("/login", handler.loginUser)

	r.POST("/create", middleware.AuthMiddleware(), handler.Create)
	r.PUT("/update", handler.Update)
	r.GET("/tweet/{id}", handler.GetById)
	r.GET("/tweets", handler.GetAllTweets)
	r.DELETE("/tweets", handler.DeleteTweetsById)

	return r
}

func (h *handler) loginUser(c *gin.Context) {
	var loginUser models.LoginUser

	err := c.BindJSON(&loginUser)
	if err != nil {
		h.logger.Errorf("cannot bind json data into struct: %v", err)
		c.JSON(400, gin.H{"error": "json data is not correct"})
		return
	}

	// Check if the username already exist in the db
	dbUser := models.LoginUser{}
	err = h.sql.QueryRow("SELECT username,password FROM users where username=$1", loginUser.UserName).Scan(&dbUser.UserName, &dbUser.Password)
	if err != nil {
		h.logger.Errorf("username doesn't exist: %v", err)
		c.JSON(400, gin.H{"error": "username doesn't  exist"})
		return
	}

	matched := auth.CheckPasswordHash(loginUser.Password, dbUser.Password)

	// Check if password matches
	if matched != true {
		h.logger.Error("username password doesn't match")
		c.JSON(400, gin.H{"error": "username password doesn't match"})
		return
	}

	// Generate a token and return
	token, err := auth.CreateToken(loginUser.UserName)
	if err != nil {
		h.logger.Error("cannot create a token")
		c.JSON(500, gin.H{"error": "cannot create a token"})
		return
	}

	// Return a session data
	c.JSON(200, gin.H{"token": token})
}

func (handler *handler) registerUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		handler.logger.Errorf("cannot bind json data into struct: %v", err)
		c.JSON(400, gin.H{"error": "json data is not correct"})
		return
	}

	// Check if the email already exist in the db
	dbUser := models.User{}
	err = handler.sql.QueryRow("SELECT first_name,last_name FROM users where email=$1", user.Email).Scan(&dbUser.FirstName, &dbUser.LastName)
	if err == nil {
		handler.logger.Errorf("email already exist")
		c.JSON(500, gin.H{"error": "email already exist"})
		return
	}

	// Check if the username already exist in the db
	err = handler.sql.QueryRow("SELECT first_name,last_name FROM users where username=$1", user.UserName).Scan(&dbUser.FirstName, &dbUser.LastName)
	if err == nil {
		handler.logger.Errorf("username already exist: %v", err)
		c.JSON(500, gin.H{"error": "username already exist"})
		return
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		handler.logger.Errorf("cannot bcrypt the password: %+v", err)
		c.JSON(500, gin.H{"error": "failed to register"})
		return
	}

	// Store the data into the database
	_, err = handler.sql.Query("INSERT INTO users (first_name,last_name, email, username, password) VALUES($1,$2,$3,$4,$5)", user.FirstName, user.LastName, user.Email, user.UserName, user.Password)
	if err != nil {
		handler.logger.Errorf("cannot insert the value to the database: %+v", err)
		c.JSON(500, gin.H{"error": "failed to register"})
		return
	}

	c.JSON(200, gin.H{"error": "successfully registered"})
}

func (handler *handler) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (handler *handler) Create(c *gin.Context) {

	// Get the users tweet and parse to the Struct
	var tweet models.Tweet

	err := c.BindJSON(&tweet)
	if err != nil {
		c.JSON(400, gin.H{"error": "json data is not correct"})
		return
	}

	userName, exist := c.Get("username")
	if !exist {
		c.JSON(500, gin.H{"error": "couldn't retrieve username"})
		return
	}

	tweet.UserName = userName.(string)
	tweet.CreateTime = time.Now()

	// TODO: Add the tweets in the database

	bx, err := json.Marshal(tweet)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot marshal"})
		return
	}

	c.JSON(200, gin.H{"data": string(bx)})

}

func (handler *handler) Update(c *gin.Context) {

}

func (handler *handler) GetById(c *gin.Context) {

}

func (handler *handler) GetAllTweets(c *gin.Context) {

}

func (handler *handler) DeleteTweetsById(c *gin.Context) {

}
