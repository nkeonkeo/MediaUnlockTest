package mediaunlocktest

import (
	"net/http"
	"strings"
)

func IqRegion(c http.Client) Result {
	resp, err := GET(c, "https://www.iq.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	s := resp.Header.Get("x-custom-client-ip")
	if s == "" {
		return Result{Status: StatusNo}
	}
	i := strings.Index(s, ":")
	if i == -1 {
		return Result{Status: StatusNo}
	}
	region := s[i+1:]
	if region == "ntw" {
		region = "tw"
	}
	return Result{Status: StatusOK, Region: region}
}
