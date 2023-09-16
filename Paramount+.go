package mediaunlocktest

import "net/http"

func ParamountPlus(c http.Client) Result {
	return Result{Status: StatusNo}
}
