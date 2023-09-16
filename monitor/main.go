package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var service bool
	var update bool
	flag.IntVar(&Interval, "interval", 60, "check interval (s)")
	flag.StringVar(&Listen, "listen", ":9101", "listen address")
	flag.BoolVar(&MUL, "mul", true, "Multination")
	flag.BoolVar(&HK, "hk", false, "Hong Kong")
	flag.BoolVar(&JP, "jp", false, "Japan")
	flag.BoolVar(&NA, "na", false, "North America")
	flag.BoolVar(&SA, "sa", false, "South")

	flag.BoolVar(&service, "service", false, "setup systemd service")
	flag.BoolVar(&update, "u", false, "check update")
	flag.Parse()
	if update {
		checkUpdate()
		return
	}
	if service {
		Service()
		return
	}
	go recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	log.Println("serve on " + Listen)
	http.ListenAndServe(Listen, nil)
}
