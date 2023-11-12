package mediaunlocktest

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

func TVer(c http.Client) Result {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://edge.api.brightcove.com/playback/v1/accounts/5102072605001/videos/ref%3Akaguyasama_01", nil)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("Accept", "application/json;pk=BCpkADawqM0_rzsjsYbC1k1wlJLU4HiAtfzjxdUmfvvLUQB-Ax6VA-p-9wOEZbCEm3u95qq2Y1CQQW1K9tPaMma9iAqUqhpISCmyXrgnlpx9soEmoVNuQpiyGsTpePGumWxSs1YoKziYB6Wz")

	resp, err := cdo(c, req)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	// log.Println(string(b))
	if strings.Contains(string(b), "error_subcode") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
	// var res struct {
	// 	ErrorSubcode string `json:"error_subcode"`
	// }
	// if err := json.Unmarshal(b, &res); err != nil {
	// 	return Result{Status: , Err: err}
	// }
}
