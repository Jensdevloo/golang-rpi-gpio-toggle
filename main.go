package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	//register routes
	http.HandleFunc("/", renderMainPage)
	http.HandleFunc("/startPitch", startPitch)
	http.HandleFunc("/startQna", startQna)
	http.HandleFunc("/nextPitch", nextPitch)
	http.HandleFunc("/nextPhase", nextPhase)

	//start webserver
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Start listenening on localhost port %v!\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
