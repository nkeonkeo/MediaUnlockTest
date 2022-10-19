package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func LineTV(c http.Client) Result {
	resp, err := GET(c, "https://www.linetv.tw/api/part/11829/eps/1/part?chocomemberId=")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res struct {
		CountryCode int
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.CountryCode == 228 {
		return Result{Success: true}
	}
	return Result{Success: false}
}
