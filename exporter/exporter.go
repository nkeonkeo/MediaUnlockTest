package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Interval  = 60
	rrcStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rrc_unblock_status",
		Help: "Region Restriction Check Status",
	}, []string{"media", "media_readable", "task"})
)

func recordMetrics() {
	go func() {
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			Check()
			for k, v := range R {
				rrcStatus.WithLabelValues(
					k,
				).Set(float64(v.Status))
			}
		}
	}()
}
