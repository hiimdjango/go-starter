package main

import (
	"fmt"
	"net/http"

	"github.com/hiimdjango/gostarter/pkg/handlers"
)

const portNumber = ":8090"

func main() {
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/", handlers.Home)

	fmt.Printf("Application listening on %s", portNumber)
	fmt.Println()
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
