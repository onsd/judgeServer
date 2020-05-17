package types

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	UserID     int    `json:"user_id"`
	QuestionID int    `json:"question_id"`
	Language   string `json:"language"`
	Answer     string `json:"answer"`
	Status     string `json:"status"`
	Result     string `json:"result"`
	Detail     string `json:"detail"`
	Error      string `json:"error"`
}

type Question struct {
	gorm.Model
	Title       string       `json:"title"`
	Body        string       `json:"body"`
	Validation  string       `json:"validation"`
	Input       string       `json:"input"`
	Output      string       `json:"output"`
	TestCases   []TestCase   `json:"testcase" gorm:"foreignkey:QuestionID"`
	SampleCases []SampleCase `json:"samplecase" gorm:"foreignkey:QuestionID"`
}

type SQSData struct {
	gorm.Model
	AnswerID  uint   `json:"answer_id"`
	Answer    string `json:"answer"`
	TestCases []TestCase
}

type TestCase struct {
	gorm.Model
	QuestionID int
	Input      string
	Output     string
}
type SampleCase struct {
	gorm.Model
	QuestionID int
	Input      string
	Output     string
}

type User struct {
	gorm.Model
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Answers []Answer `json:"questions" gorm:"foreignkey:UserID"`
}
