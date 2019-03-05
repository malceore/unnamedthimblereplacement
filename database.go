package main

import (
  "database/sql"
  "os"
  "fmt"
  "strconv"
   _ "github.com/lib/pq"
)

var (
  host     = os.Getenv("DB_HOST")
  port, _  = strconv.ParseInt(os.Getenv("DB_PORT"), 0, 64)
  user     = os.Getenv("DB_USER")
  password = os.Getenv("DB_PASS")
  dbname   = os.Getenv("DB_NAME")
  psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db, err  = sql.Open("postgres", psqlInfo)
)

func connectDatabase() {
  //psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  //db, err := sql.Open("postgres", psqlInfo)
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  //defer db.Close()
  fmt.Println("Successfully connected to Postgresql Database!")
}

func setupDatabase() {
  // Create database and table for users if non exists.
  db.Exec("CREATE DATABASE " + dbname)
  db.Exec(`CREATE TABLE users (id SERIAL PRIMARY KEY, username TEXT, password TEXT, email TEXT UNIQUE NOT NULL);`)
  fmt.Println("DEBUGG::Initial database setup complete!")
}

func registerDatabase(username string, email string, password string){
    _, err = db.Query("INSERT INTO users (username, email, password) VALUES ('" + username + "', '" + email + "', '" + password + "');")
    if err != nil {
      fmt.Println("Failed to register : " + username)
      fmt.Print(err)
    } else{
      fmt.Println("DEBUG::Registered " + username + " successfully!")
    }
}

func validateUser(username string, password string) (bool){
  entries, err := db.Query("SELECT username, password FROM users WHERE username='" + username + "' AND password='" + password + "';")
  if err != nil || entries == nil{
    return false
  }
  return true
}

func closeDatabase(){
  fmt.Println("Closing Database connection..")
  db.Close();
}
