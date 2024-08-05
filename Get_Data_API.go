package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://penerbit-staging-v2.visiku.co.id/ajukan-pendanaan")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjUyOTgiLCJpYXQiOjE3MjI2OTU0NjUsImV4cCI6MTcyMjcyNDI2NX0.KPJhKzzN5OrIcn9Cn2eUwvAnCE0eBDGMuZZGv6gWPZM"
	// response.Header.Set("Authorization", "Bearer "+token)

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
