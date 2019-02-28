package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

// The new router function creates the router and
func newRouter() *mux.Router {
	r := mux.NewRouter()
        r.HandleFunc("/", HomeHandler)
        r.HandleFunc("/home", HomeHandler)
        //r.HandleFunc("/editor", EditorHandler)
        r.HandleFunc("/login", LoginHandler).Methods("POST")
        r.HandleFunc("/logout", LogoutHandler).Methods("POST")

	staticFileDirectory := http.Dir("./res/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}
func main() {
	r := newRouter()
	http.ListenAndServe(":9191", r)
}


func HomeHandler(response http.ResponseWriter, request *http.Request){
    username := GetUserName(request)
    bytes, err := ioutil.ReadFile("res/home/index.html")
    if err != nil {
        fmt.Println("", err)
    }
    var indexbody = string(bytes)
    if len(username) > 0 {
        fmt.Fprintf(response, indexbody, username)
    } else {
	fmt.Fprintf(response, indexbody, "<a href='/login'>Login</a>")
    }
}

func EditorHandler(response http.ResponseWriter, request *http.Request){
    username := GetUserName(request)
    if len(username) > 0 {

    } else {
        http.Redirect(response, request, "/login", 302)
    }
}


// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("username")
    pass := request.FormValue("password")
    fmt.Println("input: " + name + "," + pass)
    redirectTarget := "/"
    if name == "test" && pass == "password" {
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
    http.Redirect(response, request, "/", 302)
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
/*
func GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("cookie"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}*/