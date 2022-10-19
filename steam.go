package main

import (
	"net/http"
	"strings"
)

func Steam(c http.Client) Result {
	resp, err := GET(c, "https://store.steampowered.com")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

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
