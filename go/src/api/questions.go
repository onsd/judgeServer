package api

import (
	"log"
	"net/http"

	"main/db"
	"main/types"

	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var questions []types.Question
	if err := db.Preload("TestCases").Preload("SampleCases").Order("id").Find(&questions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func GetQuestionByID(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var question types.Question

	if err := db.First(&question, c.Param("id")).Related(&question.TestCases).Related(&question.SampleCases).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, question)
}

func AddNewQuestion(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var json types.Question
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&json).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	db.Last(&json)
	c.JSON(http.StatusOK, json)
}

func UpdateQuestion(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var u types.Question
	if err := db.First(&u, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var json types.Question
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update
	u.Body = json.Body
	u.Validation = json.Validation
	u.TestCases = json.TestCases
	u.SampleCases = json.SampleCases

	db.Save(&u)
	c.JSON(http.StatusOK, u)
}

func DeleteQuestion(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var u types.Question
	if err := db.First(&u, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&u)
	c.String(204, "")
}
