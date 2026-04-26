package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "sample-ws/db"
    _"github.com/go-sql-driver/mysql" 
    "sample-ws/handlers"
    "sample-ws/middleware"
)

func main() {
    
    db.InitDB()
    defer db.DB.Close()
    
    
    router := gin.Default()


    public := router.Group("/api")
    {
        public.POST("/auth", handlers.Auth)
    }

    
    protected := router.Group("/api")
    protected.Use(middleware.ValidateJWT())
    {
        protected.GET("/query", handlers.Query)
    }

    
    log.Println("Server starting on http://localhost:8080")
    router.Run("localhost:8080")
}