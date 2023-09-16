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
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res tvbAnywhereRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.AllowInThisCountry {
		return Result{Status: StatusOK, Region: strings.ToLower(res.Country)}
	}
	return Result{Status: StatusNo}
}

type tvbAnywhereRes struct {
	AllowInThisCountry bool `json:"allow_in_this_country"`
	Country            string
}
