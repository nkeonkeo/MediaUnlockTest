package mediaunlocktest

import "net/http"

// Princess Connect Re:Dive Japan
func PCRJP(c http.Client) Result {
	resp, err := GET_Dalvik(c, "https://api-priconne-redive.cygames.jp/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 404:
		return Result{Status: StatusOK}
	case 403:
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}
