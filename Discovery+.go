package mediaunlocktest

import "net/http"

func DiscoveryPlus(c http.Client) Result {
	return Result{Status: StatusNo}
}
