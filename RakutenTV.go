package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func RakutenTV(c http.Client) Result {
    resp, err := c.Get("https://rakuten.tv")
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    
    if err != nil {
		return Result{Status: StatusFailed}
	}
	
	
	if strings.Contains(bodyString, "waitforit") {
		return Result{Status: StatusNo}
	}
	
	if resp.StatusCode == 200 {
    	return Result{Status: StatusOK}
	}
	
	return Result{Status: StatusUnexpected}
}