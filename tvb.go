package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func TVBAnywhere(c http.Client) Result {
	resp, err := GET(c, "https://uapisfm.tvbanywhere.com.sg/geoip/check/platform/android")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

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
