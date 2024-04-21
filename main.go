package main

import (
	"fmt"
	authcontroller "github.com/jeypc/go-auth/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)

	fmt.Println("Server jalan di port 3000")
	http.ListenAndServe(":3000", nil)
}
