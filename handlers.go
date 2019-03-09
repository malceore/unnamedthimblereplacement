package main

import (
	"net/http"
        "github.com/gorilla/securecookie"
	"io/ioutil"
	"fmt"
)

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

/*
** Start of handler functions
*/
func HomeHandler(response http.ResponseWriter, request *http.Request){
    username := GetUserName(request)
    bytes, err := ioutil.ReadFile("res/home/index.html")
    if err != nil {
        fmt.Println("", err)
    }
    var indexbody = string(bytes)
    if len(username) > 0 {
        fmt.Fprintf(response, indexbody, `Welcome ` + username + `! <form action="/create"> <input name="username" type="hidden" value="` + username + `"><input type="submit" value="New Project"> </form>`)
    } else {
        var login = `Please <a href="/login">Login</a> or <a href="/register">Register</a> `
        fmt.Fprintf(response, indexbody, login)
    }
}

// for POST
func RegisterHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("username")
    pass1 := request.FormValue("password")
//    pass2 := request.FormValue("retype")
    email := request.FormValue("email")

    fmt.Println("DEBUG::Registering: " + name + "," + email)

    redirectTarget := "/"
    if name != "" && password != "" && email != "" {
           SetCookie(name, response)
           redirectTarget = "/home"
           registerDatabase(name, email, pass1)
    } else {
            redirectTarget = "/register"
    }
    http.Redirect(response, request, redirectTarget, 302)
}

// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("username")
    pass := request.FormValue("password")
    redirectTarget := "/"
    if validateUser(name, pass) {
           SetCookie(name, response)
           redirectTarget = "/home"
    } else {
            redirectTarget = "/login"
    }
    http.Redirect(response, request, redirectTarget, 302)
}

// for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
    cookie := &http.Cookie{
        Name:   "cookie",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
    http.Redirect(response, request, "/home", 302)
}


// for POST
func CreateProjectHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("username")
    fmt.Println("DEBUG::Create new project For: " + name)
    newProject(name)
    http.Redirect(response, request, "/editor", 302)
}


/*
** Start of assisting functions.
*/
func GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("cookie"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}

func SetCookie(userName string, response http.ResponseWriter) {
    value := map[string]string{
        "name": userName,
    }
    if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
        cookie := &http.Cookie{
            Name:  "cookie",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(response, cookie)
    }
}
