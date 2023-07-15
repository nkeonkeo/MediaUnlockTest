package mediaunlocktest

import "net/http"

func BritBox(c http.Client) Result {
	return Result{Success: false}
}
