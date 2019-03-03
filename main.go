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
        //r.HandleFunc("/editor", EditorHandler)
        r.HandleFunc("/login", LoginHandler).Methods("POST")
        r.HandleFunc("/logout", LogoutHandler).Methods("POST")

	staticFileDirectory := http.Dir("./res/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}
func main() {
        connect()
	r := newRouter()
	http.ListenAndServe(":9191", r)
}

