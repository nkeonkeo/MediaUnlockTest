package mediaunlocktest

import "net/http"

func SlingTV(c http.Client) Result {
	return Result{Success: false}
}
