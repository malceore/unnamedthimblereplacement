package main

import (
  "database/sql"
  "fmt"
   _ "github.com/lib/pq"
)

const (
  host     = "192.168.0.107"
  port     = 5432
  user     = "test"
  password = "test"
  dbname   = "app"
)


func connect() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  fmt.Println("Successfully connected to Postgresql Database!")
}

func setup(){
  sqlStatment := '
    CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
    );'

  _, err = db.Exec(sqlStatement)
  /*if err != nil {
    panic(err)
  }*/

  fmt.Println("Initial database setup complete!")
}