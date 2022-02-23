package controller

import (
	"fmt"
	"log"
	"net/http"
)

func InitializeServerRoutes() {
	http.HandleFunc("/conversion", ConversionHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
