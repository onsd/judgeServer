package db

import (
	"log"
	"main/types"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		"host="+os.Getenv("HOSTNAME")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("USER")+
			" sslmode=disable"+
			" dbname="+os.Getenv("DBNAME")+
			" password="+os.Getenv("PASSWORD"),
	)
	if err != nil {
		log.Printf("Error at open Database: %v", err)
		return nil, err
	}

	return db, nil
}

func InitDB(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		log.Printf("Error at getDB()\n %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&types.Answer{}, &types.TestCase{}, &types.SampleCase{}, &types.Question{}, &types.User{})
}
