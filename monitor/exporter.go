package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Interval = 60
	Listen   = ":9101"

	rrcStatus *prometheus.GaugeVec
)

func init() {
	rrcStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "media_unblock_status",
		Help: "Media Unlock Status",
	}, []string{"mediaName"})
}

func update() {
	Check()
	for k, v := range R {
		// log.Println(k, v)
		rrcStatus.WithLabelValues(
			k,
		).Set(float64(v.Status))
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
