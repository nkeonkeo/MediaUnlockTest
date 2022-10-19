package main

import (
	"net/http"
)

func J_COM_ON_DEMAND(c http.Client) Result {
	c.CheckRedirect = nil
	resp, err := GET(c, "https://linkvod.myjcom.jp/auth/login")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		return Result{Success: false}
	case 502:
		return Result{Success: false}
	}
	return Result{Success: true}
}
