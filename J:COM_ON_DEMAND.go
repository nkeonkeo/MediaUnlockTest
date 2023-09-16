package mediaunlocktest

import (
	"net/http"
)

func J_COM_ON_DEMAND(c http.Client) Result {
	c.CheckRedirect = nil
	resp, err := GET(c, "https://linkvod.myjcom.jp/auth/login")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		return Result{Status: StatusNo}
	case 502:
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
