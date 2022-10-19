package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func Dazn(c http.Client) Result {
	resp, err := PostJson(c, "https://startup.core.indazn.com/misl/v5/Startup",
		`{"LandingPageKey":"generic","Languages":"zh-CN,zh,en","Platform":"web","PlatformAttributes":{},"Manufacturer":"","PromoCode":"","Version":"2"}`,
	)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	// log.Println(string(b))
	var res daznRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Region.IsAllowed {
		return Result{
			Success: true,
			Region:  res.Region.GeolocatedCountry,
		}
	}
	return Result{
		Success: false,
		Info:    res.Region.DisallowedReason,
	}
}

type daznRegion struct {
	IsAllowed             bool
	DisallowedReason      string
	GeolocatedCountry     string
	GeolocatedCountryName string
}

type daznRes struct {
	Region daznRegion
}
