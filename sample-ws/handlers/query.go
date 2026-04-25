package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "sample-ws/db"
    "sample-ws/models"
)

func Query(c *gin.Context) {
    
    usernameInterface, exists := c.Get("username")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
        return
    }

    username := usernameInterface.(string)
     
    
    var userData models.QueryResponse
    userData.User = username

    query := "SELECT data FROM user_data WHERE username = ? ORDER BY created_at DESC LIMIT 1"
    err := db.DB.QueryRow(query, username).Scan(&userData.Data)

    if err != nil {
        
        userData.Message = "Welcome to the secure endpoint!"
        userData.Data = "No specific data found for this user. Please update your profile."
    } else {
        userData.Message = "Here is your secure information"
    }

    c.JSON(http.StatusOK, userData)
}