package main

import (
	"fmt"
	"net/http"
)

func main() {
	//router := gin.Router()
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Sprintln("hehe")
	//})
	if err := http.ListenAndServe(":7077", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
