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

// Forgot to give table a filename value, so dumb.
var file struct {
    fileId string
    contents string
}

func connectDatabase() {
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  //defer db.Close()
  fmt.Println("DB::Successfully connected to Postgresql Database!")
}

func setupDatabase() {
  // Create database and table for users if non exists.
  _,err = db.Exec("CREATE DATABASE IF NOT EXIST " + dbname + ";")
  _,err = db.Query("SELECT userId FROM users WHERE username=user")
  if (err != nil) {
    // construct tables
    db.Exec(`CREATE TABLE users (userId SERIAL PRIMARY KEY, username TEXT UNIQUE, password TEXT, email TEXT UNIQUE NOT NULL);`)
    db.Exec(`CREATE TABLE projects (projectId SERIAL PRIMARY KEY, userId SERIAL REFERENCES users(userId), name TEXT);`)
    db.Exec(`CREATE TABLE files (fileId SERIAL PRIMARY KEY, projectId SERIAL REFERENCES projects(projectId), fileName TEXT, contents TEXT);`)

    var entries *sql.Rows
    // Inject default user and project
    entries, err = db.Query("INSERT INTO users (username, email, password) VALUES ('user', 'default@defaulters.com', 'user') RETURNING userId;")
    var userId string=returnSingleRow(entries)

    entries, err=db.Query("INSERT INTO projects (userId, name) VALUES ('" + userId +  "', 'Default') RETURNING projectId;")
    var projectId string=returnSingleRow(entries)

    _, err=db.Query("INSERT INTO files (projectId, fileName, contents) VALUES ('" + projectId + ", 'index.html', 'PGgxPldlbGNvbWUgdG8gVGhpbWJsZTwvaDE+');")
    fmt.Println("DB::Initial database setup complete!")
  }
}

func registerDatabase(username string, email string, password string){
    _, err = db.Query("INSERT INTO users (username, email, password) VALUES ('" + username + "', '" + email + "', '" + password + "');")
    if err != nil {
      fmt.Println("DEBUG::DB::Failed to register : " + username)
      fmt.Print(err)
    } else{
      fmt.Println("DEBUG::DB::Registered " + username + " successfully!")
    }
}

func validateUser(username string, password string) (bool){
  entries, err := db.Query("SELECT username, password FROM users WHERE username='" + username + "' AND password='" + password + "';")
  if err != nil || entries == nil{
    return false
  }
  return true
}

func getFiles(projectId string) (*sql.Rows){
  entries, err := db.Query("SELECT fileId, contents FROM files WHERE projectId='" + projectId + "';")
  if err != nil || entries == nil{
    return nil
  }
  return entries
}

func saveFile(fileId string, contents string) {
  _, err := db.Query("UPDATE files SET contents = '" + contents + "' WHERE fileId = " + fileId + ";")
  if err != nil{
    fmt.Println(err)
  }
}


func cloneProject(targetProjectId string, userId string) (string) {
  // First we add a new project for the requested user.
  // Learned about returning, which sends back the variable after creation!
  var newProjectId string
  results, _ := db.Query("INSERT INTO projects (userId, name) VALUES ('" + userId + "', 'Default') RETURNING projectId;")
  for results.Next() {
    err := results.Scan(&newProjectId)
    fmt.Println(err)
  }

  var contents string
  var filename string
  // Next we select and iterate all the files in the old project to add them to the new.
  entries, _ := db.Query("SELECT contents, filename FROM files WHERE projectId='" + targetProjectId + "';")
  for entries.Next() {
    entries.Scan(&contents, &filename)
    _, err :=db.Query("INSERT INTO files (contents, projectId, filename) VALUES ('" + contents + "', '" + newProjectId + "', '" + filename + "') RETURNING fileId;")
    fmt.Println(err)
  }

  fmt.Println("DEBUG::DB: Cloning project for " + userId + "... ")
  return newProjectId
}


func getUserId(username string) (string){
  var contents string
  // Next we select and iterate all the files in the old project to add them to the new.
  entries, _ := db.Query("SELECT userid FROM users WHERE username='" + username + "';")
  for entries.Next() {
    entries.Scan(&contents)
    //_, err :=db.Query("INSERT INTO files (contents, projectId) VALUES ('" + contents + "', '" + newProjectId + "') RETURNING fileId;")
    fmt.Println(err)
  }
  return contents
}

func returnSingleRow(rows *sql.Rows) (string){
  var contents string
  for rows.Next() {
    rows.Scan(&contents)
  }
  return contents
}

func closeDatabase(){
  fmt.Println("Closing Database connection..")
  db.Close();
}
