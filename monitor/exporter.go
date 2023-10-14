package main

import (
	"strings"
	"sync"

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
	updateLock  = sync.Mutex{}
)

func init() {
	rrcStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "media_unblock_status",
		Help: "Media Unlock Status",
	}, []string{"node", "mediaName", "region"})
}

func update() {
	updateLock.Lock()
	defer updateLock.Unlock()

	Check()

	rrcStatus.Reset()
	for _, v := range R {
		rrcStatus.WithLabelValues(
			Node,
			v.Name,
			strings.ToUpper(v.Value.Region),
		).Set(float64(v.Value.Status))
	}
}

func recordMetrics() {
	go update()
	gocron.Every(Interval).Seconds().Do(update)
	<-gocron.Start()
}
