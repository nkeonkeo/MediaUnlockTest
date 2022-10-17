package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func DMM(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://api-p.videomarket.jp/v3/api/play/keyauth?playKey=4c9e93baa7ca1fc0b63ccf418275afc2&deviceType=3&bitRate=0&loginFlag=0&connType=", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("User-Agent", UA_Browser)
	req.Header.Add("X-Authorization", "2bCf81eLJWOnHuqg6nNaPZJWfnuniPTKz9GXv5IS")

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res dmmRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Status.Code == 403 {
		return Result{Success: false}
	}
	// if res.Status.Code == 408 {
	// 	return Result{Success: true}
	// }
	return Result{Success: true}
}

type dmmRes struct {
	Status struct {
		Code    int
		Message string
	}
}
