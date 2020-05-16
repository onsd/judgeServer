package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/sqs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Answer struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
	Status     string    `json:"status"`
	Result     string    `json:"result"`
	Detail     string    `json:"detail"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// r.GET("/answers/:id/status", getAnswerStatus)
func getAnswerStatus(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var answer Answer
	if err := db.First(&answer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNotImplemented, answer)
}

// r.POST("/answers/:id", addNewAnswer)
func addNewAnswer(c *gin.Context) {
	sqs := sqs.NewSQS("python")
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var postedJSON Answer
	if err := c.ShouldBindJSON(&postedJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	questionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postedJSON.QuestionID = questionID
	if err := db.Create(&postedJSON).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	jsonBytes, err := json.Marshal(postedJSON)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	sqs.SendMessage(string(jsonBytes))
	db.Last(&postedJSON)
	c.JSON(http.StatusOK, postedJSON)
}

func getAnswers(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var answers []Answer
	if err := db.Order("id").Find(&answers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, answers)
}
func getAnswerByID(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var answer Answer
	if err := db.First(&answer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, answer)
}

func updateAnswer(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var u Answer
	if err := db.First(&u, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var postedJSON Answer
	if err := c.BindJSON(&postedJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update
	u.Status = "Evalated"
	u.Result = postedJSON.Result
	u.Detail = postedJSON.Detail
	fmt.Println(u)
	db.Save(&u)
	c.JSON(http.StatusOK, u)
}
