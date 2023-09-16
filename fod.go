package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func FOD(c http.Client) Result {
	resp, err := GET(c, "https://geocontrol1.stream.ne.jp/fod-geo/check.xml?time=1624504256")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "true") {
		return Result{Status: StatusOK}
	}
	if strings.Contains(s, "false") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}
