package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func TVBAnywhere(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://uapisfm.tvbanywhere.com.sg/geoip/check/platform/android", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res tvbAnywhereRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.AllowInThisCountry {
		return Result{Success: true, Region: strings.ToLower(res.Country)}
	}
	return Result{Success: false}
}

type tvbAnywhereRes struct {
	AllowInThisCountry bool `json:"allow_in_this_country"`
	Country            string
}
