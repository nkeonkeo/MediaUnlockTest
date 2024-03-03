package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func FXNOW(c http.Client) Result {
	resp, err := GET(c, "https://fxnow.fxnetworks.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if strings.Contains(string(b), "is not accessible") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
