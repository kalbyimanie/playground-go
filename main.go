package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
)

var counter int = 0

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	counter++
	counterConv := strconv.Itoa(counter)
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip":   GetIP(r),
		"hits": counterConv,
	})

	w.Write(resp)

}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func handleRequest(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaultRoute)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	port := os.Getenv("PORT_NUMBER")
	if len(port) == 0 {
		fmt.Println("required env `PORT_NUMBER`")
	} else {
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("Server starting on port %s....", port)
		fmt.Println()
		handleRequest(port)
	}
}
