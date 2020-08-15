package main

import (
	"flag"
	"log"

	"github.com/thejasbabu/k6-ping-pong/ping"
	"github.com/thejasbabu/k6-ping-pong/pong"
)

func main() {
	var isPing, isPong bool
	var port int
	var endpoint string
	flag.BoolVar(&isPing, "ping", false, "run as ping service")
	flag.BoolVar(&isPong, "pong", false, "run as pong service")
	flag.IntVar(&port, "port", 8080, "port which will be used to listen")
	flag.StringVar(&endpoint, "endpoint", "", "endpoint of pong service")
	flag.Parse()

	if isPing && endpoint == "" {
		log.Fatal("pong endpoint not specified")
	} else if isPing {
		pinger := ping.NewPinger(endpoint, port)
		pinger.Start()
	} else if isPong {
		ponger := pong.NewPonger(port)
		ponger.Start()
	} else {
		flag.PrintDefaults()
		log.Fatal("specify to run as ping or pong")
	}
}
