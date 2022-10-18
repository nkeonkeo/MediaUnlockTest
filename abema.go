package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func Abema(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://api.abema.io/v1/ip/check?device=android", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Dalvik)

	var resp *http.Response
	for i := 0; i < 3; i++ {
		if resp, err = c.Do(req); err == nil {
			break
		}
	}
	if err != nil {
		return Result{Success: false, Err: err}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	// log.Println(string(b))
	var res abemaRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.IsoCountryCode == "JP" {
		return Result{Success: true}
	}
	return Result{Success: false, Info: "Oversea Only"}
}

type abemaRes struct {
	IsoCountryCode string
}
