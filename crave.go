package mediaunlocktest

import "net/http"

func Crave(c http.Client) Result {
	return Result{Success: false}
}
