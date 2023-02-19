package models

import (
	"backend/db_access"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID                 uint64 `gorm:"primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	FirstName          string `json:"firstname" binding:"required"`
	LastName           string `json:"lastname" binding:"required"`
	Email              string `json:"email" binding:"required"`
	Picture            string `json:"picture" binding:"required"`
	Duration           string `json:"duration"`
	DBName             string `json:"dbname"`
	IsTempAccessActive bool
}

type Request struct {
	User User `json:"user" binding:"required"`
}

type ActiveRequest struct {
	Email    string `json:"email" binding:"required"`
	Duration string `json:"duration"`
	Action   bool   `json:"action"`
	DBName   string `json:"dbname"`
}

var ErrUserExists = errors.New("user already exists")

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func DeleteExpiredUsers(watsonDB *gorm.DB, db *sql.DB) {
	var users []User
	err := watsonDB.Where(&User{IsTempAccessActive: true}).Find(&users).Error
	if err != nil {
		log.Println("DeleteExpiredUsers: error during job")
	}

	now := time.Now()
	log.Printf("starting user cleanup for DEV access")
	for _, user := range users {
		duration, _ := time.ParseDuration(user.Duration)
		if now.Sub(user.UpdatedAt) > duration {
			log.Printf("removing user - '%s'", user.Email)
			activeRequest := ActiveRequest{Email: user.Email, Duration: user.Duration, Action: false}
			setActiveStatusForUser(watsonDB, activeRequest, &user)

			devUser := db_access.DevDbUser{Username: user.Email, Duration: user.Duration}
			devUser.DeleteUserFromDB(db)
		}
	}
	log.Println("active user cleanup complete")
}

func setActiveStatusForUser(db *gorm.DB, activeRequest ActiveRequest, tempUser *User) {
	tempUser.IsTempAccessActive = activeRequest.Action
	tempUser.Duration = activeRequest.Duration
	tempUser.DBName = activeRequest.DBName
	db.Save(&tempUser)
}

func SetTempActiveStatusForUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var activeRequest ActiveRequest
		var tempUser User

		if err := c.ShouldBindJSON(&activeRequest); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		err := db.Where(&User{Email: activeRequest.Email}).First(&tempUser).Error
		if err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			return
		}

		setActiveStatusForUser(db, activeRequest, &tempUser)

		c.JSON(200, gin.H{"user": &tempUser})
	}
}

func GetAllTempActiveUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		err := db.Where(&User{IsTempAccessActive: true}).Find(&users).Error
		if err == nil {
			c.JSON(200, gin.H{"users": &users})
			return
		}
		c.JSON(404, gin.H{"message": err.Error()})
	}
}

func checkIfUserExists(db *gorm.DB, user User) error {
	var returnUser User
	err := db.Where(&User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}).First(&returnUser).Error
	if err != nil {
		return nil
	}
	return ErrUserExists
}

func GetUserWithEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		var user User
		err := db.Where(&User{Email: email}).First(&user).Error
		if err == nil {
			c.JSON(200, gin.H{"user": &user})
		} else {
			c.JSON(404, gin.H{"message": err.Error()})
		}
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseUint(s, 0, 19)
	return err == nil
}

func GetUserWithID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		if !isNumeric(param) {
			c.JSON(400, gin.H{"message": "id is not number"})
			return
		}

		idNumber, _ := strconv.ParseUint(c.Param("id"), 0, 19)
		var user User
		err := db.Where(&User{ID: idNumber}).First(&user).Error
		if err == nil {
			c.JSON(200, gin.H{"user": &user})
		} else {
			c.JSON(404, gin.H{"message": err.Error()})
		}
	}
}

func AddUserToWatson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		user := request.User

		if checkIfUserExists(db, user) == ErrUserExists {
			log.Println(ErrUserExists)
			c.JSON(409, gin.H{"message": ErrUserExists.Error()})
			return
		}

		res := db.Create(&user)
		if res.Error != nil {
			log.Println(res.Error)
			c.JSON(500, gin.H{"message": res.Error.Error()})
			return
		}

		c.JSON(200, gin.H{"message": fmt.Sprintf("user '%s' created in watson", user.Email)})
	}
}

func DeleteUserFromWatson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		requestedUser := request.User

		if checkIfUserExists(db, requestedUser) != ErrUserExists {
			c.JSON(404, gin.H{"message": "user does not exist in watson"})
			return
		}

		var user User
		getUserForDeletion := db.Where(&User{Email: requestedUser.Email}).First(&user)
		if getUserForDeletion.Error != nil {
			log.Println(getUserForDeletion.Error)
			c.JSON(500, gin.H{"message": getUserForDeletion.Error.Error()})
			return
		}

		res := db.Unscoped().Delete(&User{}, user.ID)
		if res.Error != nil {
			log.Println(res.Error)
			c.JSON(500, gin.H{"message": res.Error.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "user deleted in watson"})
	}
}
