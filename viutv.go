package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func ViuTV(c http.Client) Result {
	resp, err := PostJson(c, "https://api.viu.now.com/p8/3/getLiveURL",
		`{"callerReferenceNo":"20210726112323","contentId":"099","contentType":"Channel","channelno":"099","mode":"prod","deviceId":"29b3cb117a635d5b56","deviceType":"ANDROID_WEB"}`,
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
		return Result{Status: StatusErr, Err: err}
	}
	if res.ResponseCode == "SUCCESS" {
		return Result{Status: StatusOK}
	} else if res.ResponseCode == "GEO_CHECK_FAIL" {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}
