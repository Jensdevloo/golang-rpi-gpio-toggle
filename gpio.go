package main

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func toggleLeds() {
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
