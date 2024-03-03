package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func ParamountPlus(c http.Client) Result {
	resp, err := GET(c, "https://www.paramountplus.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "intl") {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
