package main

import (
	"log"
	"net/http"
)

func renderMainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(html))
}

func startPitch(w http.ResponseWriter, r *http.Request) {
	log.Println("Pitch started!")
	message := "Start pitch!"
	w.Write([]byte(message))
}

func startQna(w http.ResponseWriter, r *http.Request) {
	log.Println("QnA started!")
	message := "Start QnA!"
	w.Write([]byte(message))
}

func nextPitch(w http.ResponseWriter, r *http.Request) {
	log.Println("Move to next pitch name!")
	message := "Move to next pitch name!"
	w.Write([]byte(message))
}

func nextPhase(w http.ResponseWriter, r *http.Request) {
	log.Println("Move to next phase led!")
	message := "Move to next phase led!"
	w.Write([]byte(message))
}
