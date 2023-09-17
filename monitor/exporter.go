package main

import (
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Interval    = 60
	Listen      = ":9101"
	ChangeIpCmd string

	rrcStatus *prometheus.GaugeVec
)

func init() {
	rrcStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "media_unblock_status",
		Help: "Media Unlock Status",
	}, []string{"mediaName", "region"})
}

func update() {
	r := make([]*result, len(R))
	copy(r, R)

	Check()

	for _, v := range r {
		rrcStatus.DeleteLabelValues(
			v.Name,
			strings.ToUpper(v.Value.Region),
		)
	}
	for _, v := range R {
		rrcStatus.WithLabelValues(
			v.Name,
			strings.ToUpper(v.Value.Region),
		).Set(float64(v.Value.Status))
	}
}

func recordMetrics() {
	update()
	t := time.NewTicker(time.Duration(Interval) * time.Second)
	defer t.Stop()
	for range t.C {
		update()
	}
}
