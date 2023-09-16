package mediaunlocktest

import "net/http"

func Funimation(c http.Client) Result {
	return Result{Status: StatusNo}
}
