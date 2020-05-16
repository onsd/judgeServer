package main

import (
	"os"
	"regexp"

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

	r := gin.New()

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
	r.PUT("/answers/:id", updateAnswer)
	r.GET("/answers/:id", getAnswerByID)

	r.Run(":" + os.Getenv("PORT"))
}
