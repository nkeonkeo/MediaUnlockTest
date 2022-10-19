package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func NowE(c http.Client) Result {
	r, err := http.NewRequest("POST", "https://webtvapi.nowe.com/16/1/getVodURL", strings.NewReader(
		`{"contentId":"202105121370235","contentType":"Vod","pin":"","deviceId":"W-60b8d30a-9294-d251-617b-c12f9d0c","deviceType":"WEB"}`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	r.Header.Set("User-Agent", UA_Browser)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(r)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res noweRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.ResponseCode == "PRODUCT_INFORMATION_INCOMPLETE" {
		return Result{Success: true}
	} else if res.ResponseCode == "GEO_CHECK_FAIL" {
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}

type noweRes struct {
	ResponseCode string
}
