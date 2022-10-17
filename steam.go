package main

import (
	"net/http"
	"strings"
)

func Steam(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://store.steampowered.com", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}

	for _, c := range resp.Cookies() {
		if c.Name == "steamCountry" {
			i := strings.Index(c.Value, "%")
			if i == -1 {
				return Result{Success: false}
			}
			return Result{Success: true, Region: strings.ToLower(c.Value[:i])}
		}
	}
	return Result{Success: false}
}
