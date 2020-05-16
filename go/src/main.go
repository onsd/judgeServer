package main

import (
	"os"
	"regexp"

	"main/api"
	"main/db"
	"main/index"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	r := gin.Default()

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	r.Use(logger.SetLogger())

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Logger()

	r.Use(logger.SetLogger(logger.Config{
		Logger:         &subLog,
		UTC:            true,
		SkipPath:       []string{"/skip"},
		SkipPathRegexp: rxURL,
	}))
	r.LoadHTMLGlob("templates/*")

	r.GET("/", ping) // 挨拶
	r.GET("/init", db.InitDB)
	main := r.Group("/")
	{
		main.GET("/index", index.GetIndex)
		main.GET("/questions/:id", index.GetQuestionByID)
		main.GET("/submit/:id", index.GetSubmitForm)
		main.POST("/submit/:id", index.PostSubmit)
		main.GET("/answers/:id/status", index.GetAnswerStatus)
		main.GET("/submitquestion", index.GetQuestionSubmitForm)

	}
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/users/:id", api.GetUserByID)   //指定した id の user を表示
		apiGroup.POST("/users", api.AddNewUser)       //user を追加
		apiGroup.PUT("/users/:id", api.UpdateUser)    //指定した id の user を更新
		apiGroup.DELETE("/users/:id", api.DeleteUser) //指定した id の user をapiGroup
		apiGroup.GET("/questions", api.GetQuestions)
		apiGroup.GET("/questions/:id", api.GetQuestionByID)
		apiGroup.POST("/questions", api.AddNewQuestion)
		apiGroup.PUT("/questions/:id", api.UpdateQuestion)
		apiGroup.DELETE("/questions/:id", api.DeleteQuestion)
		apiGroup.GET("/answers", api.GetAnswers)
		apiGroup.GET("/answers/:id/status", api.GetAnswerStatus)
		apiGroup.POST("/answers/:id", api.AddNewAnswer)
		apiGroup.PUT("/answers/:id", api.UpdateAnswer)
		apiGroup.GET("/answers/:id", api.GetAnswerByID)
	}

	r.Run(":" + os.Getenv("PORT"))
}
