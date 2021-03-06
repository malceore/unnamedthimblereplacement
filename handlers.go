package main

import (
	"net/http"
        "github.com/gorilla/securecookie"
	"io/ioutil"
	"fmt"
        "strings"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

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
        fmt.Fprintf(response, indexbody, `Welcome ` + username + `! <form action="/remix/1"> <input name="username" type="hidden" value="` + username + `"><input type="submit" value="New Project"> </form>`)
    } else {
        var login = `Please <a href="/login">Login</a> or <a href="/register">Register</a> `
        fmt.Fprintf(response, indexbody, login)
    }
}

func RegisterHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("username")
    pass1 := request.FormValue("password")
    email := request.FormValue("email")
    fmt.Println("HA::Registering: " + name + "," + email)

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

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
    cookie := &http.Cookie{
        Name:   "thimble_cookie",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
    http.Redirect(response, request, "/home", 302)
}


func RemixProjectHandler(response http.ResponseWriter, request *http.Request) {
    var project = strings.Split(request.URL.Path, "/")
    if project[2] == "" {
      fmt.Fprintf(response, "HA::Error Occured, please specify a project!")
    } else{
      name := request.FormValue("username")
      fmt.Println("HA::Remix project For: " + name)
      var newProjectId string = cloneProject(project[2], getUserId(name))
      http.Redirect(response, request, "/editor/" + newProjectId, 302)
   }
}


func SaveHandler(response http.ResponseWriter, request *http.Request){
     fileId := request.FormValue("fileid")
     contents := request.FormValue("contents")
     saveFile(fileId, contents)
     //fmt.Println("DEBUG::Saved file " + contents + " " + fileId)
}


func EditorHandler(response http.ResponseWriter, request *http.Request){
    // Protected Websites requires login cookies.
    fmt.Println("DEBUG::Editor Handler")
    var exists = CheckCookie(request);
    if (!exists){
      fmt.Println("Not allowed!")
      http.Redirect(response, request, "/login", 302)
    }

    // Confirm that a project was given or just throw an error.
    var project = strings.Split(request.URL.Path, "/")
    //fmt.Println(project)
    if project[2] == "" {
      fmt.Fprintf(response, "Error Occured, please specify a project!")
    } else{
     // If it does then we need to return the contents of it's file sin a javascript array, injected into the index.html page.
      bytes, err := ioutil.ReadFile("res/editor/index.html")
      if err != nil {
          fmt.Println(err)
      }
      var indexbody = string(bytes)

      fmt.Println("DEBUG::Getting files for this project " + project[2])
      var files = getFiles(project[2])
      var contents = "["
      if files != nil{
        indexbody = string(bytes)
        var fileId string
        var fileContents string
        notLast := files.Next()
        for notLast {
          err := files.Scan(&fileId, &fileContents)
          if err != nil {
            fmt.Println(err)
          } else {
            // Fense post problem.
            contents += " '" + fileId + "', '" + fileContents + "' "
            notLast = files.Next()
            if notLast {
              contents += ","
            }
          }
        }
      }
      contents+="]"
      fmt.Fprintf(response, indexbody, contents)
    }
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
            Name:  "cookiethimble",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(response, cookie)
    }
}

func CheckCookie(r *http.Request) (exists bool)  {
	_, err := r.Cookie("cookiethimble")
	if err != nil {
		fmt.Println("cookiethimble")
		return false
	}
        return true
}