package main

import (
	"log"
	"net/http"

	"wordora/app/routes"
	"wordora/databases"
	"wordora/databases/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer db.Close()

	log.Println("🚀 Running database migrations...")
	migrations.MigrateDatabase(db)

	// r := gin.Default()

	// router := routes.SetupRouter(db)
	router := routes.SetupRouter(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "✅ Server is running on localhost:8080!"})
	})

	log.Println("🚀 Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}

	// router.Run(":8080")
}
