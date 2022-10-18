package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func LiTV(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://www.litv.tv/vod/ajax/getUrl", strings.NewReader(
		`{"type":"noauth","assetId":"vod44868-010001M001_800K","puid":"6bc49a81-aad2-425c-8124-5b16e9e01337"}`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("Content-Type", "application/json")

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
		ErrorMessage interface{}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.ErrorMessage == nil {
		return Result{Success: true}
	}
	return Result{}
}
