package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("main\n")

	router := httprouter.New()

	// router
	router.GET("/user", getUser)
	router.POST("/user", createUser)

	// サーバー起動
	if err := http.ListenAndServe(":5050", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
