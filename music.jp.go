package mediaunlocktest

import (
	"io"
	"net/http"
)

func MusicJP(c http.Client) Result {
	resp, err := GET(c, "https://overseaauth.music-book.jp/globalIpcheck.js")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if string(b) == "" {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
