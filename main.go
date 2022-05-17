package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func versionedHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page of peng service")
	fmt.Println("Endpoint: homePage[peng]")
}

func pong(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://pong-container:8082/response")

	if err != nil {
		fmt.Fprintf(w, "peng cannot reach pong service")
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Fprintf(w, "peng ... "+string(body)+"\n")
	fmt.Printf("Endpoint: peng tries to reach pong service\n")
}

func ping(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://ping-container:8081/response")

	if err != nil {
		fmt.Fprintf(w, "peng cannot reach ping service")
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Fprintf(w, "peng ... "+string(body)+"\n")
	fmt.Printf("Endpoint: peng tries to access ping service\n")
}

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Has reached peng service - version 1")
}

func handleRequests() {
	http.HandleFunc("/", versionedHomePage)
	http.HandleFunc("/peng-to-pong", pong)
	http.HandleFunc("/peng-to-ping", ping)
	http.HandleFunc("/response", response)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func main() {
	fmt.Println("Starting rest server...")
	handleRequests()
}
