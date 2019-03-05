package main

import (
	//"os"
        //"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// The new router function creates the router and
func newRouter() *mux.Router {
	r := mux.NewRouter()
        r.HandleFunc("/", HomeHandler)
        r.HandleFunc("/home", HomeHandler)
        r.HandleFunc("/login", LoginHandler).Methods("POST")
        r.HandleFunc("/logout", LogoutHandler).Methods("POST")
        r.HandleFunc("/register", RegisterHandler).Methods("POST")
	staticFileDirectory := http.Dir("./res/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
        //fmt.Println("DEBUG::printing ENV " + os.Environ())
        connectDatabase()
        // Will need logic to check if it's populated first.
        setupDatabase()
	r := newRouter()
	http.ListenAndServe(":9191", r)
}

