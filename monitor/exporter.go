package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Interval    uint64 = 60
	Listen             = ":9101"
	ChangeIpCmd string
	Node        = ""
	rrcStatus   *prometheus.GaugeVec
)

func init() {
	rrcStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "media_unblock_status",
		Help: "Media Unlock Status",
	}, []string{"node", "mediaName", "region"})
}

func update() {
	T := NewTest()
	T.Check()
	rrcStatus.Reset()
	log.Println("update:")
	for _, v := range T.Results {
		fmt.Println(v.Name, v.Value.Region, v.Value.Status)
		rrcStatus.WithLabelValues(
			Node,
			v.Name,
			strings.ToUpper(v.Value.Region),
		).Set(float64(v.Value.Status))
	}
}

func recordMetrics() {
	go update()
	t := time.NewTicker(time.Duration(Interval) * time.Second)
	for range t.C {
		go update()
	}
}
