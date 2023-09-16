package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func Starz(c http.Client) Result {
	resp, err := GET(c, "https://www.starz.com/sapi/header/v1/starz/us/09b397fc9eb64d5080687fc8a218775b", H{"Referer", "https://www.starz.com/us/en/"})
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	authorization := string(b)
	resp2, err := GET(c, "https://auth.starz.com/api/v4/User/geolocation", H{"AuthTokenAuthorization", authorization})
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	b2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res struct {
		IsAllowedAccess  bool
		IsAllowedCountry bool
		IsKnownProxy     bool
		Country          string
	}
	if err := json.Unmarshal(b2, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.IsAllowedAccess && res.IsAllowedCountry && !res.IsKnownProxy {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
