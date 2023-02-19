package main

import (
	"backend/db_access"
	"backend/middlewares"
	"backend/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	s := gocron.NewScheduler(time.UTC)

	// DB Connection
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DBURL := fmt.Sprintf("%s:%s@tcp(%s)/", DbUser, DbPassword, DbHost)

	// Connect to MGMT db
	MGMT_DB := "watson.db"
	mgmtDb, err := gorm.Open(sqlite.Open(MGMT_DB), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	models.MigrateDB(mgmtDb)

	// Connect to DB
	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// CORS config
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token"}

	// Gin config
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(corsConfig))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ok"})
	})

	api := router.Group("/api")
	{
		api.POST("/db-access", middlewares.CheckIfHasToken(), db_access.CreateUser(db))
		api.GET("/db-access", middlewares.CheckIfHasToken(), db_access.GetAllDBs(db))
		api.DELETE("/db-access", middlewares.CheckIfHasToken(), db_access.DeleteUser(db))

		api.POST("/user", middlewares.CheckIfHasToken(), models.AddUserToWatson(mgmtDb))
		api.POST("/user/access-status", middlewares.CheckIfHasToken(), models.SetTempActiveStatusForUser(mgmtDb))
		api.GET("/user/:email", middlewares.CheckIfHasToken(), models.GetUserWithEmail(mgmtDb))
		api.GET("/user/id/:id", middlewares.CheckIfHasToken(), models.GetUserWithID(mgmtDb))
		api.GET("/user/all-temp-active", middlewares.CheckIfHasToken(), models.GetAllTempActiveUsers(mgmtDb))
		api.DELETE("/user", middlewares.CheckIfHasToken(), models.DeleteUserFromWatson(mgmtDb))
	}

	s.Every(60).Seconds().Do(models.DeleteExpiredUsers, mgmtDb, db)

	s.StartAsync()
	router.Run(":8081") // listen and serve on 0.0.0.0:8081
}
