package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Answer struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// r.GET("/answers/:id/status", getAnswerStatus)
func getAnswetStatus(c *gin.Context) {

}

// r.POST("/answers/:id", addNewAnswer)
func addNewAnswer(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()
}
