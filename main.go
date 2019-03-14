package main

import (
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
	r.HandleFunc("/create", CreateProjectHandler)
        r.HandleFunc("/editor/{id:[0-9]+}", EditorHandler)
        r.HandleFunc("/project/{id:[0-9]+}", ProjectHandler)
	staticFileDirectory := http.Dir("./res/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
        connectDatabase()
        setupDatabase()
	r := newRouter()
	http.ListenAndServe(":9191", r)
	closeDatabase()
}

