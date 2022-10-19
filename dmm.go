package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func DMM(c http.Client) Result {
	resp, err := GET(c, "https://api-p.videomarket.jp/v3/api/play/keyauth?playKey=4c9e93baa7ca1fc0b63ccf418275afc2&deviceType=3&bitRate=0&loginFlag=0&connType=")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
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
