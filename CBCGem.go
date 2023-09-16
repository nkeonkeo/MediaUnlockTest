package mediaunlocktest

import "net/http"

func CBCGem(c http.Client) Result {
	return Result{Status: StatusNo}
}
