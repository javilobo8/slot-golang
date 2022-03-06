package main

import (
	"flag"
	"math/rand"
	"time"

	app "slot-golang/pkg"
	"slot-golang/rtp"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	mode := flag.String("mode", "server", "run mode, server or rtp")
	flag.Parse()

	switch *mode {
	case "rtp":
		rtp.CalculateRTP()
	case "server":
		app.CreateApp()
	}
}
