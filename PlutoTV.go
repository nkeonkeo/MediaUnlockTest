package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func PlutoTV(c http.Client) Result {
	resp, err := GET(c, "https://pluto.tv/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "thanks-for-watching") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
