package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", ping)                   // 挨拶
	r.GET("/users/:id", getUserByID)   //指定した id の user を表示
	r.POST("/users", addNewUser)       //user を追加
	r.PUT("/users/:id", updateUser)    //指定した id の user を更新
	r.DELETE("/users/:id", deleteUser) //指定した id の user を削除

	r.GET("/questions", getQuestions)
	r.GET("/questions/:id", getQuestionByID)
	r.POST("/questions", addNewQuestion)
	r.PUT("/questions/:id", updateQuestion)
	r.DELETE("/questions/:id", deleteQuestion)

	r.GET("/answers", getAnswers)
	r.GET("/answers/:id/status", getAnswerStatus)
	r.POST("/answers/:id", addNewAnswer)

	r.Run(":" + os.Getenv("PORT"))
}
