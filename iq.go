package main

import (
	"net/http"
	"strings"
)

func IqRegion(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://www.iq.com", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Browser)
	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
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
