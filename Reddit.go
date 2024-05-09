package mediaunlocktest

import (
    "io"
	"net/http"
	"strings"
)

func Reddit(c http.Client) Result {
	resp, err := http.Get("https://www.reddit.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    
    if err != nil {
		return Result{Status: StatusFailed}
	}
	
	if resp.StatusCode == 200 {
		return Result{Status: StatusOK}
	}
	
	if resp.StatusCode == 403 && strings.Contains(bodyString, "blocked") {
		return Result{Status: StatusNo}
	}

	return Result{Status: StatusUnexpected}
}