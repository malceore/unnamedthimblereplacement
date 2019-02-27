package main

import (
	"net/http"
)

func GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("cookie"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}

