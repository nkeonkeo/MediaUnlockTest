package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func HboGoAisa(c http.Client) Result {
	resp, err := GET(c, "https://api2.hbogoasia.com/v1/geog?lang=undefined&version=0&bundleId=www.hbogoasia.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res hboRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Territory == "" {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK, Region: strings.ToLower(res.Country)}
}

type hboRes struct {
	Country   string
	Territory string
}
