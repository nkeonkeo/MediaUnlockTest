package mediaunlocktest

import (
    "io"
    "strings"
	"net/http"
)

func BBCiPlayer(c http.Client) Result {
    resp, err := http.Get("https://open.live.bbc.co.uk/mediaselector/6/select/version/2.0/mediaset/pc/vpid/bbc_one_london/format/json/jsfunc/JS_callbacks0")
    
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
	    if strings.Contains(bodyString, "geolocation") {
	    	return Result{Status: StatusNo}
    	}
		return Result{Status: StatusOK}
	}

	return Result{Status: StatusFailed}
}