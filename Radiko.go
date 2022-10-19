package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func Radiko(c http.Client) Result {
	resp, err := GET(c, "https://radiko.jp/area?_=1625406539531")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `classs="OUT"`) {
		return Result{Success: false}
	}
	if strings.Contains(s, "JAPAN") {
		return Result{Success: true}
	}
	return Result{Success: false}
}
