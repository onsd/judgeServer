package api

import (
	"encoding/json"
	"fmt"
	"log"
	"main/db"
	"main/sqs"
	"main/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// r.GET("/answers/:id/status", getAnswerStatus)
func GetAnswerStatus(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var answer types.Answer
	if err := db.First(&answer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNotImplemented, answer)
}

// r.POST("/answers/:id", addNewAnswer)
func AddNewAnswer(c *gin.Context) {
	sqs := sqs.NewSQS("python")
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var postedJSON types.Answer
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

func GetAnswers(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var answers []types.Answer
	if err := db.Order("id").Find(&answers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, answers)
}
func GetAnswerByID(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var answer types.Answer
	if err := db.First(&answer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, answer)
}

func UpdateAnswer(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var u types.Answer
	if err := db.First(&u, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var postedJSON types.Answer
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
