package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!",
	})
}
