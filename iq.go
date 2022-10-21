package mediaunlocktest

import (
	"net/http"
	"strings"
)

func IqRegion(c http.Client) Result {
	resp, err := GET(c, "https://www.iq.com")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	s := resp.Header.Get("x-custom-client-ip")
	if s == "" {
		return Result{Success: false}
	}
	i := strings.Index(s, ":")
	if i == -1 {
		return Result{Success: false}
	}
	region := s[i+1:]
	if region == "ntw" {
		region = "tw"
	}
	return Result{Success: true, Region: region}
}
