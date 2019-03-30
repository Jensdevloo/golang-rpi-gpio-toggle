package main

import (
	"fmt"
	"time"
 	"net/http"
 	"strings"

	"github.com/stianeikeland/go-rpio"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))

	fmt.Println("opening gpio")
	
	toggleLeds()

}

func toggleLeds(){
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	pin := rpio.Pin(18)
	pin.Output()

	defer rpio.Close()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	toggleLeds()
}
