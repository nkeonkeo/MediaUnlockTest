package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func Niconico(c http.Client) Result {
	resp, err := GET(c, "https://www.nicovideo.jp/watch/so40278367")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}

	if strings.Contains(string(b), "同じ地域") {
		return Result{Success: false}
	}
	return Result{Success: true}
}
