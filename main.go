package main

import (
	"log"
	"net/http"

	"wordora/app/routes"
	"wordora/databases"
	"wordora/databases/migrations"

	_ "wordora/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Wordora Blogs API
// @version 1.0
// @description This is a sample server for Wordora.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer db.Close()

	// log.Println("🚀 Running database migrations...")
	migrations.MigrateDatabase(db)
	// migrations.RollbackDatabase(db)

	router := routes.SetupRouter(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Server is running on port localhost:8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
