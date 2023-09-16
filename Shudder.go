package mediaunlocktest

import "net/http"

func Shudder(c http.Client) Result {
	return Result{Status: StatusNo}
}
