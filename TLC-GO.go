package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func TlcGo(c http.Client) Result {
	resp, err := GET(c, "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location/dnsfeed")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if strings.Contains(string(b), `"country":"US"`) {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
