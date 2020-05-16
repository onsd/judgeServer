package index

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

func GetIndex(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var questions []types.Question
	if err := db.Order("id").Find(&questions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"questions": questions,
	})
}

func GetQuestionByID(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var question types.Question
	if err := db.Preload("SampleCases").First(&question, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(question.SampleCases)
	c.HTML(http.StatusOK, "question.html", gin.H{
		"question":    question,
		"sampleCases": question.SampleCases,
	})
}

func GetSubmitForm(c *gin.Context) {
	c.HTML(http.StatusOK, "submit.html", gin.H{
		"id": 1,
	})
}
func GetQuestionSubmitForm(c *gin.Context) {
	c.HTML(http.StatusOK, "submitQuestion.html", gin.H{})
}

func PostSubmit(c *gin.Context) {
	id := c.Param("id")
	questionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Request.ParseForm()
	language := c.Request.Form["language"]
	answer := c.Request.Form["answer"]
	if language[0] == "" || answer[0] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqs := sqs.NewSQS("python")
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()

	var postedJSON types.Answer
	postedJSON.Answer = answer[0]
	postedJSON.Language = language[0]
	postedJSON.QuestionID = questionID
	postedJSON.Status = "SUBMIT"

	if err := db.Create(&postedJSON).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	var question types.Question

	if err := db.First(&question, c.Param("id")).Related(&question.TestCases).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	sqsData := types.SQSData{
		AnswerID:  postedJSON.ID,
		Answer:    postedJSON.Answer,
		TestCases: question.TestCases,
	}
	jsonBytes, err := json.Marshal(sqsData)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	sqs.SendMessage(string(jsonBytes))
	db.Last(&postedJSON)
	url := "/answers/" + strconv.Itoa(int(sqsData.AnswerID)) + "/status"
	c.Redirect(302, url)
}

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

	c.HTML(http.StatusOK, "result.html", gin.H{
		"answer": answer,
	})
}
