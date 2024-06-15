package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Listening on :2621")
	http.ListenAndServe(":2621", nil)
}
