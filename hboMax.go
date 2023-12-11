package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func HBOMax(c http.Client) Result {
	resp, err := GET(c, "https://www.hbomax.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	s := string(b)
	if strings.Contains(s, "geo-availability") {
		return Result{Status: StatusNo}
	}

	return Result{Status: StatusNo}
}
