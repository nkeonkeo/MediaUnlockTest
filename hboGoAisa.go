package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func HboGoAisa(c http.Client) Result {
	resp, err := GET(c, "https://api2.hbogoasia.com/v1/geog?lang=undefined&version=0&bundleId=www.hbogoasia.com")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res hboRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Territory == "" {
		return Result{Success: false}
	}
	return Result{Success: true, Region: strings.ToLower(res.Country)}
}

type hboRes struct {
	Country   string
	Territory string
}
