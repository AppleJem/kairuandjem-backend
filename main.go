package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []user{
	{ID: `1`, Username: `user1`, Password: `password1`},
	{ID: `2`, Username: `user2`, Password: `password2`},
	{ID: `3`, Username: `user3`, Password: `password3`},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", registerUser)

	router.Run("0.0.0.0:8080")
}

func registerUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
