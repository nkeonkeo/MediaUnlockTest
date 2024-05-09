package mediaunlocktest

import (
    "io"
	"net/http"
	"strings"
)

func WikipediaEditable(c http.Client) Result {
	resp, err := http.Get("https://zh.wikipedia.org/w/index.php?title=Wikipedia%3A%E6%B2%99%E7%9B%92&action=edit")
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    
    if err != nil {
		return Result{Status: StatusFailed}
	}
	
	if strings.Contains(bodyString, "Banned") {
		return Result{Status: StatusNo}
	}
	
	if resp.StatusCode == 200 {
		return Result{Status: StatusOK}
	}
	
	return Result{Status: StatusUnexpected}
}