package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var service bool
	var update bool
	var version bool
	flag.IntVar(&Interval, "interval", 60, "check interval (s)")
	flag.StringVar(&Listen, "listen", ":9101", "listen address")
	flag.StringVar(&Node, "node", "", "node")
	flag.BoolVar(&MUL, "mul", true, "Multination")
	flag.BoolVar(&HK, "hk", false, "Hong Kong")
	flag.BoolVar(&TW, "tw", false, "Taiwan")
	flag.BoolVar(&JP, "jp", false, "Japan")
	flag.BoolVar(&NA, "na", false, "North America")
	flag.BoolVar(&SA, "sa", false, "South")

	flag.BoolVar(&service, "service", false, "setup systemd service")
	flag.BoolVar(&update, "u", false, "check update")
	flag.BoolVar(&version, "v", false, "show version")

	flag.Parse()

	if version {
		fmt.Println("unlock-monitor "+Version, "("+runtime.GOOS+"_"+runtime.GOARCH+" "+runtime.Version()+" build-time: "+buildTime+")")
		return
	}
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
