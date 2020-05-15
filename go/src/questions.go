package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Questions は問題ページの問題を示します
type Questions struct {
	ID         int       `json:"id"`
	Body       string    `json:"body"`
	Validation string    `json:"validation"`
	Input      string    `json:"input"`
	Output     string    `json:"output"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func getQuestions(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var questions []Questions
	if err := db.Order("id").Find(&questions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func getQuestionByID(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var question Questions
	if err := db.First(&question, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, question)
}
