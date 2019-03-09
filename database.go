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
  db.Exec("CREATE DATABASE " + dbname + ";")
  // On next iteration username should be unique..
  db.Exec(`CREATE TABLE users (userId SERIAL PRIMARY KEY, username TEXT, password TEXT, email TEXT UNIQUE NOT NULL);`)
  db.Exec(`CREATE TABLE projects (projectId SERIAL PRIMARY KEY, userId SERIAL REFERENCES users(userId), name TEXT);`)
  db.Exec(`CREATE TABLE files (fileId SERIAL PRIMARY KEY, projectId SERIAL REFERENCES projects(projectId), contents TEXT);`)
  //db.Query("INSERT INTO users (username, email, password) VALUES ('default', 'default@defaulters.com', 'test');")
  // Not yet sure how we get the user ID
  //entries, err := db.Query("SELECT userId FROM users WHERE username='default';")
  //db.Query("INSERT INTO projects (username, email, password) VALUES ('default', 'default@defaulters.com', 'test');")
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

/*
** Fetches the user id for the username given, calls clone using it and the default user/project. returns
*/
func newProject(username string) (bool){
  var userId string
  entries, _ := db.Query("SELECT userId FROM users WHERE username='" + username + "';")
  for entries.Next() {
    err := entries.Scan(&userId)
    if err != nil {
      fmt.Println("DEBUG:: Fatal no entries!")
    } else {
      cloneProject("1", "1", userId);
    }
  }
  return true
}

func cloneProject(targetId string, oldProjectId string, newId string){
  // First we add a new project for the requested user.
  // Learned about returning, which sends back the variable after creation!
  var newProjectId string
  results, _ := db.Query("INSERT INTO projects (userId, name) VALUES ('" + newId + "', 'Default') RETURNING projectId;")
  for results.Next() {
    err := results.Scan(&newProjectId)
    fmt.Println(err)
  }

  var contents string
  // Next we select and iterate all the files in the old project to add them to the new.
  entries, _ := db.Query("SELECT contents FROM files WHERE projectId='" + oldProjectId + "';")
  for entries.Next() {
    entries.Scan(&contents)
    _, err :=db.Query("INSERT INTO files (contents, projectId) VALUES ('" + contents + "', '" + newProjectId + "') RETURNING fileId;")
    fmt.Println(err)
  }

  fmt.Println("DEBUG:: Cloning default project for " + newId + "... ")
}


func closeDatabase(){
  fmt.Println("Closing Database connection..")
  db.Close();
}
