package mediaunlocktest

import (
	"net/http"
)

func DirecTVGO(c http.Client) Result {
	resp, err := http.Get("https://www.directvgo.com/registrarse")
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	defer resp.Body.Close()
    
    if err != nil {
		return Result{Status: StatusFailed}
	}
	
	if resp.StatusCode == 403 {
	    return Result{Status: StatusNo}
	}
	
	if resp.StatusCode == 200 {
		return Result{Status: StatusOK}
	}
	
	return Result{Status: StatusFailed}
}