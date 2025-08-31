package main

import (
	"context"
	"errors"
	"log"
	"manager_project/config"
	"manager_project/config/migrations"
	"manager_project/middleware"
	"manager_project/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}
	server := gin.Default()

	server.Use(middleware.RateLimitMiddleware())

	router.SetupAllRoutes(server, db)

	server.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "database connection error",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "test database ping failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
		})
	})

	srv := &http.Server{
		Addr:    ":8989",
		Handler: server,
	}

	go func() {
		log.Printf("server listen at %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	} else {
		log.Println("Server exiting gracefully")
	}
}
