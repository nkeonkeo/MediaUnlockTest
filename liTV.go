package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func LiTV(c http.Client) Result {
	resp, err := PostJson(c, "https://www.litv.tv/vod/ajax/getUrl",
		`{"type":"noauth","assetId":"vod44868-010001M001_800K","puid":"6bc49a81-aad2-425c-8124-5b16e9e01337"}`,
	)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		ErrorMessage interface{}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.ErrorMessage == nil {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
