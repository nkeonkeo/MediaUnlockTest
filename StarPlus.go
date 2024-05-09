package mediaunlocktest

import (
    "net/http"
    "io/ioutil"
    "regexp"
    "strings"
)


func SupportStarPlus(loc string) bool {
	var STARPLUS_SUPPORT_COUNTRY = []string{
        "BR", "MX", "AR", "CL", "CO", "PE", "UY", "EC", "PA", "CR", "PY", "BO", "GT", "NI", "DO", "SV", "HN", "VE",
    }
	for _, s := range STARPLUS_SUPPORT_COUNTRY {
		if loc == s {
			return true
		}
	}
	return false
}

func StarPlus(c http.Client) Result {

	resp, err := http.Get("https://www.starplus.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	
    if err != nil {
        return Result{Status: StatusFailed}
    }

	if resp.StatusCode == 403 {
		return Result{Status: StatusBanned}
	}
	
	if resp.StatusCode == 200 {
	    re := regexp.MustCompile(`Region:\s+([A-Za-z]{2})`)
        matches := re.FindStringSubmatch(string(body))
        if len(matches) >= 2 {
            if SupportStarPlus(matches[1]) {
                return Result{Status: StatusOK, Region: strings.ToLower(matches[1])}
            }
            return Result{Status: StatusNo}
	    }
		return Result{Status: StatusUnexpected}
	}

	return Result{Status: StatusFailed}
}