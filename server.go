package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/readiness", Readiness)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "I am %s and I am %s years old", name, age)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s Password: %s", user, pass)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("cats/cats.txt")

	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}

	fmt.Fprintf(w, "Cats: %s", string(data))
}

// func Healthz(w http.ResponseWriter, r *http.Request) {
// 	duration := time.Since(startedAt)

// 	if duration.Seconds() > 25 {
// 		w.WriteHeader(500)
// 		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
// 	} else {
// 		w.WriteHeader(200)
// 		w.Write([]byte("OK"))
// 	}
// }

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 15 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}
}
