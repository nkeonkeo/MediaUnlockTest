package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func Paravi(c http.Client) Result {
	resp, err := PostJson(c, "https://api.paravi.jp/api/v1/playback/auth",
		`{"meta_id":17414,"vuid":"3b64a775a4e38d90cc43ea4c7214702b","device_code":1,"app_id":1}`,
	)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res struct {
		Error struct {
			Type string
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Error.Type == "Forbidden" {
		return Result{Status: StatusNo}
	}
	if res.Error.Type == "Unauthorized" {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusOK}
}
