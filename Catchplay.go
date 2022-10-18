package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func Catchplay(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://sunapi.catchplay.com/geo", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("authorization", "Basic NTQ3MzM0NDgtYTU3Yi00MjU2LWE4MTEtMzdlYzNkNjJmM2E0Ok90QzR3elJRR2hLQ01sSDc2VEoy")

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res struct {
		Code string
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Code == "100016" {
		return Result{Success: false}
	}
	return Result{Success: true}
}
