package main

import (
	"fmt"
	"log"

	"budgetly-backend/internal/config"
	"budgetly-backend/internal/database"
	"budgetly-backend/internal/transaction"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := database.Connect(connString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(nil)

	router := gin.Default()

	transactionRepo := transaction.NewRepository(db)

	transactionService := transaction.NewService(
		transactionRepo,
	)

	transactionHandler := transaction.NewHandler(
		transactionService,
	)

	api := router.Group("/api")

	transactionGroup := api.Group("/transactions")

	transactionHandler.RegisterRoutes(
		transactionGroup,
	)

	router.GET(
		"/health",
		func(c *gin.Context) {
			c.JSON(
				200,
				gin.H{
					"status": "ok",
				},
			)
		},
	)

	log.Printf("server running on :%s", cfg.AppPort)

	router.Run(":" + cfg.AppPort)
}