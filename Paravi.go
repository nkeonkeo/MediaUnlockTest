package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Paravi(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://api.paravi.jp/api/v1/playback/auth", strings.NewReader(
		`{"meta_id":17414,"vuid":"3b64a775a4e38d90cc43ea4c7214702b","device_code":1,"app_id":1}`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", UA_Browser)

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
		Error struct {
			Type string
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Error.Type == "Forbidden" {
		return Result{Success: false}
	}
	if res.Error.Type == "Unauthorized" {
		return Result{Success: true}
	}
	return Result{Success: false, Info: "failed"}
}
