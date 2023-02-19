package db_access

import (
	"backend/slackclient"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var ErrNoDBWithThatName = errors.New("no db with that name")
var ErrUserAlreadyExists = errors.New("a user with name exists already")

type DevDbUser struct {
	Username string `json:"username" binding:"required"`
	Duration string `json:"duration" binding:"required"`
	DBName   string `json:"db"`
	Password string
}

// Generate password for the user
func (u DevDbUser) generatePassword() string {
	rand.Seed(time.Now().UTC().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz0123456789#"
	strlen := 16
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// Delete user from DB
func (u DevDbUser) DeleteUserFromDB(db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf("DROP USER '%s'@'10.%%'", u.Username))
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("DeleteUser: '%s' - %w (you can't delete a user that has not been created yet.)", u.Username, err)
	}

	_, err = db.Exec("FLUSH PRIVILEGES")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Printf("deleted user: '%s'", u.Username)

	slackPreText := fmt.Sprintf("%s - user deleted.", u.Username)
	slackclient.SendMessage(slackPreText)

	return nil
}

// check if the db name is valid or in the RDS
func checkIfDbExists(db *sql.DB, dbname string) error {
	var name string
	rows, err := db.Query("show databases")
	if err != nil {
		log.Println(err)
		return fmt.Errorf("checkIfDbExists: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return fmt.Errorf("checkIfDbExists: %w", err)
		}
		checkName := dbname
		if name == checkName {
			return nil // db name is correct. exit the function.
		}
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("checkIfDbExists: %w", err)
	}
	return ErrNoDBWithThatName
}

// check if user exists
func checkIfUserExists(db *sql.DB, username string) error {
	var name string
	rows, err := db.Query("SELECT user FROM mysql.user")
	if err != nil {
		log.Println(err)
		return fmt.Errorf("checkIfUserExists: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return fmt.Errorf("checkIfUserExists: %w", err)
		}
		if name == username {
			return ErrUserAlreadyExists
		}
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("checkIfUserExists: %w", err)
	}
	return nil
}

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user DevDbUser
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := checkIfUserExists(db, user.Username); err != nil {
			errMsg := fmt.Errorf("CreateUser: '%s' - %w", user.Username, err)
			log.Println(errMsg.Error())
			c.JSON(409, gin.H{"message": errMsg.Error()})
			return
		}

		// check if user didnt enter db name
		if user.DBName == "" {
			c.JSON(400, gin.H{"message": "must specify which db you want to access"})
			return
		}

		// check if duration is valid
		d, err := time.ParseDuration(user.Duration)
		if err != nil {
			errMsg := fmt.Errorf("CreateUser: %w", err)
			log.Println(errMsg.Error())
			c.JSON(400, gin.H{"message": errMsg.Error()})
			return
		}

		if d > 8*time.Hour {
			errMsg := fmt.Errorf("CreateUser: Duration must be under 8 hours")
			log.Println(errMsg.Error())
			c.JSON(400, gin.H{"message": errMsg.Error()})
			return
		}

		user.Password = user.generatePassword()

		if err := checkIfDbExists(db, user.DBName); err != nil {
			errMsg := fmt.Errorf("CreateUser: '%s' - %w", user.DBName, err)
			log.Println(errMsg.Error())
			c.JSON(404, gin.H{"message": errMsg.Error()})
			return
		}

		// CREATE USER
		_, err = db.Exec(fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'10.%%' IDENTIFIED BY '%s'", user.Username, user.Password))
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		// GRANT PRIVILEGES
		_, err = db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'10.%%'", user.DBName, user.Username))
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		// FLUSH PRIVILEGES
		_, err = db.Exec("FLUSH PRIVILEGES")
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		slackPreText := fmt.Sprintf("%s - user created at %s DB. User will be deleted in %s.", user.Username, user.DBName, user.Duration)
		slackclient.SendMessage(slackPreText)

		c.JSON(200, gin.H{"message": "user created", "user": user})
	}
}

func GetAllDBs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		sysDbs := []string{"information_schema",
			"mysql",
			"performance_schema",
			"sys"}

		var dbName string
		var dbList []string
		rows, err := db.Query("show databases")
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		defer rows.Close()
		for rows.Next() {
			isAppend := true
			err := rows.Scan(&dbName)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{"message": err.Error()})
				return
			}
			for _, element := range sysDbs {
				if element == dbName && isAppend {
					isAppend = false
				}
			}
			if isAppend {
				dbList = append(dbList, dbName)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"db-list": dbList})
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user DevDbUser
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		if err := checkIfUserExists(db, user.Username); err == nil {
			errMsg := fmt.Errorf("DeleteUser: '%s' does not exist", user.Username)
			log.Println(errMsg.Error())
			c.JSON(404, gin.H{"message": errMsg.Error()})
			return
		}

		if err := user.DeleteUserFromDB(db); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "user deleted"})
	}
}
