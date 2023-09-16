package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func HamiVideo(c http.Client) Result {
	resp, err := GET(c, "https://hamivideo.hinet.net/api/play.do?id=OTT_VOD_0000249064&freeProduct=1")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res struct {
		Code string
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Code == "06001-107" {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
