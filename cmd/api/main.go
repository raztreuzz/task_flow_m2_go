package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"task_flow_m2_go/internal/platform/config"
	"task_flow_m2_go/internal/platform/database"
)

func loadEnvName() string {
	data, err := os.ReadFile(".APP_ENV")
	if err != nil {
		return "dev"
	}
	return strings.TrimSpace(string(data))
}

func main() {
	env := loadEnvName()
	envFile := fmt.Sprintf("env/.env.%s", env)

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("error loading env file: %s", envFile)
	}

	cfg := config.Load()

	db := database.NewMySQL(cfg.MySQLDSN)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"env":    env,
		})
	})

	err = r.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal(err)
	}
}
