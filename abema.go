package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func Abema(c http.Client) Result {
	resp, err := GET_Dalvik(c, "https://api.abema.io/v1/ip/check?device=android")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res abemaRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.IsoCountryCode == "JP" {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusRestricted, Info: "Oversea Only"}
}

type abemaRes struct {
	IsoCountryCode string
}
