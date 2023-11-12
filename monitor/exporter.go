package main

import (
	"strings"

	"github.com/jasonlvhit/gocron"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Interval    uint64 = 60
	Listen             = ":9101"
	ChangeIpCmd string
	Node        = ""
	rrcStatus   *prometheus.GaugeVec
	updateting  = false
)

func init() {
	rrcStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "media_unblock_status",
		Help: "Media Unlock Status",
	}, []string{"node", "mediaName", "region"})
}

func update() {

	if updateting {
		return
	}
	updateting = true

	Check()

	rrcStatus.Reset()
	for _, v := range R {
		rrcStatus.WithLabelValues(
			Node,
			v.Name,
			strings.ToUpper(v.Value.Region),
		).Set(float64(v.Value.Status))
	}
	updateting = false
}

func recordMetrics() {
	go update()
	if Interval%60 == 0 {
		gocron.Every(Interval / 60).Minute().Do(update)
	} else {
		gocron.Every(Interval).Seconds().Do(update)
	}
	<-gocron.Start()
}
