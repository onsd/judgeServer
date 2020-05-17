package api

import (
	"log"
	"net/http"

	"main/db"
	"main/types"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var users []types.User
	if err := db.Order("id").Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var user types.User
	if err := db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func AddNewUser(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var json types.User
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

func UpdateUser(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var u types.User
	if err := db.First(&u, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var json types.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update
	u.Name = json.Name
	u.Email = json.Email
	db.Save(&u)
	c.JSON(http.StatusOK, u)
}

func DeleteUser(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Printf("Error at db.GetDB()\n %v", err)
	}
	defer db.Close()

	var u types.User
	if err := db.First(&u, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&u)
	c.String(204, "")
}
