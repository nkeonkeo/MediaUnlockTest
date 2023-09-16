package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func TlcGo(c http.Client) Result {
	resp, err := GET(c, "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	resp2, err := GET(c, "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	b2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if strings.Contains(string(b), `"country":"US"`) && strings.Contains(string(b2), `"country":"US"`) {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
