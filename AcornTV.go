package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func AcornTV(c http.Client) Result {
	resp, err := GET(c, "https://acorn.tv/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return Result{Status: StatusBanned}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "Not yet available in your country") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
