package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func KKTV(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://api.kktv.me/v3/ipcheck", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		Data struct {
			Country   string
			IsAllowed bool `json:"is_allowed"`
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Data.Country == "TW" && res.Data.IsAllowed {
		return Result{Success: true}
	}
	return Result{Success: false}
}
