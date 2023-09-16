package mediaunlocktest

import "net/http"

func PeacockTV(c http.Client) Result {
	return Result{Status: StatusNo}
}
