package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func BritBox(c http.Client) Result {
	resp, err := GET(c, "https://www.britbox.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "locationnotsupported") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
