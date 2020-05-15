package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", ping) // 挨拶
	// r.GET("/questions", getUsers)          // user の一覧を表示
	r.GET("/users/:id", getUserByID)   //指定した id の user を表示
	r.POST("/users", addNewUser)       //user を追加
	r.PUT("/users/:id", updateUser)    //指定した id の user を更新
	r.DELETE("/users/:id", deleteUser) //指定した id の user を削除

	r.GET("/questions/:id", getQuestionByID)
	r.GET("/questions", getQuestions)

	r.Run(":" + os.Getenv("PORT"))
}
