package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func NowE(c http.Client) Result {
	resp, err := PostJson(c, "https://webtvapi.nowe.com/16/1/getVodURL",
		`{"contentId":"202105121370235","contentType":"Vod","pin":"","deviceId":"W-60b8d30a-9294-d251-617b-c12f9d0c","deviceType":"WEB"}`,
	)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res noweRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusUnexpected, Err: err}
	}
	if res.ResponseCode == "PRODUCT_INFORMATION_INCOMPLETE" {
		return Result{Status: StatusOK}
	} else if res.ResponseCode == "GEO_CHECK_FAIL" {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}

type noweRes struct {
	ResponseCode string
}
