package main

import (
	"net/http"
	"strings"
)

func ViuCom(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://www.viu.com", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	if location := resp.Header.Get("location"); location != "" {
		region := strings.Split(location, "/")[4]
		if region == "no-service" {
			return Result{Success: false}
		}
		return Result{Success: true, Region: region}
	}
	return Result{Success: false}
}
