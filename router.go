package main

import (
	"learning/api/handlers"
	"net/http"
)

func initializeRouter() {
	http.HandleFunc("/", handlers.Home)
}
