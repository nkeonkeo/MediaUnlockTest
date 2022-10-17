package main

import (
	"io"
	"net/http"
	"strings"
)

func PrimeVideo(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://www.primevideo.com", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Browser)
	c.CheckRedirect = nil
	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	s := string(b)

	if i := strings.Index(s, `"currentTerritory":`); i != -1 {
		return Result{
			Success: true,
			Region:  s[i+20 : i+22],
		}
	}
	return Result{Success: false}
}
