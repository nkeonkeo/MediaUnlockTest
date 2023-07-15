package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func FXNOW(c http.Client) Result {
	resp, err := GET(c, "https://fxnow.fxnetworks.com")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	return Result{Success: !strings.Contains(string(b), "is not accessible")}
}
