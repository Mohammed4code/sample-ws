package handlers

import (
	"net/http"
	"time"

	"sample-ws/db"
	"sample-ws/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func Auth(c *gin.Context) {
    var creds models.Credentials

    if err := c.BindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

   //استدعا الدالة هنا
    dbUser, err := db.GetUserByCredentials(creds.Username, creds.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  dbUser.ID,
        "username": dbUser.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
        "iat":      time.Now().Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token":   tokenString,
        "message": "Authentication successful",
        "user":    dbUser.Username,
    })
}