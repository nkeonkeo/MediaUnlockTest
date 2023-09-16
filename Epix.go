package mediaunlocktest

import "net/http"

func Epix(c http.Client) Result {
	return Result{Status: StatusNo}
}
