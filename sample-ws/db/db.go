package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    
    DB, err = sql.Open("mysql", "root:mysql@tcp(localhost:3306)/sample_ws?parseTime=true")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    fmt.Println("Connected to MySQL database")

    
    createTables()
}

func createTables() {
    userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := DB.Exec(userTable)
    if err != nil {
        log.Fatal("Failed to create users table:", err)
    }

    queryTable := `
    CREATE TABLE IF NOT EXISTS user_data (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL,
        data TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (username) REFERENCES users(username)
    )`

    _, err = DB.Exec(queryTable)
    if err != nil {
        log.Fatal("Failed to create user_data table:", err)
    }


    insertSampleUsers()
}

func insertSampleUsers() {
    users := []struct {
        username string
        password string
    }{
        {"mohammed", "123mod"},
        {"ahmed", "321ahm"},
        {"hassan", "147has"},
        {"ali", "351ali"},
    }

    for _, user := range users {
        _, err := DB.Exec("INSERT IGNORE INTO users (username, password) VALUES (?, ?)", 
            user.username, user.password)
        if err != nil {
            log.Printf("Warning: Failed to insert user %s: %v\n", user.username, err)
        }
    }

    
    sampleData := []struct {
        username string
        data     string
    }{
        {"mohammed", "Mohammed works for an IT company"},
        {"ahmed", "Ahmed studies at the university"},
        {"hassan", "Hassan works as a software engineer."},
        {"ali", "Ali has experience with Go and JavaScript."},
    }

    for _, data := range sampleData {
        _, err := DB.Exec("INSERT IGNORE INTO user_data (username, data) VALUES (?, ?)",
            data.username, data.data)
        if err != nil {
            log.Printf("Warning: Failed to insert data for %s: %v\n", data.username, err)
        }
    }
}