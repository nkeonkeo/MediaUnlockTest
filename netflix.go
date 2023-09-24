package mediaunlocktest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func NetflixRegion(c http.Client) Result {
	// 70143836 绝命毒师
	// 80018499 test
	// 81280792 乐高
	resp, err := GET(c, "https://www.netflix.com/title/81280792")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp2, err := GET(c, "https://www.netflix.com/title/70143836")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	_, err = io.ReadAll(resp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 404 && resp2.StatusCode == 404 {
		return Result{Status: StatusRestricted, Info: "Originals Only"}
	}
	if resp.StatusCode == 403 && resp2.StatusCode == 403 {
		return Result{Status: StatusBanned}
	}
	if (resp.StatusCode == 200 || resp.StatusCode == 301) || (resp2.StatusCode == 200 || resp2.StatusCode == 301) {
		resp3, err := GET(c, "https://www.netflix.com/title/80018499")
		if err != nil {
			return Result{Status: StatusNetworkErr, Err: err}
		}
		defer resp3.Body.Close()
		_, err = io.ReadAll(resp3.Body)
		if err != nil {
			log.Fatal(err)
		}
		u := resp3.Header.Get("location")
		if u == "" {
			return Result{Status: StatusOK, Region: "us"}
		}
		// log.Println("nf", u)
		t := strings.SplitN(u, "/", 5)
		if len(t) < 5 {
			return Result{Status: StatusUnexpected}
		}
		return Result{Status: StatusOK, Region: strings.SplitN(t[3], "-", 2)[0]}
	}
	return Result{Status: StatusUnexpected}
}

func NetflixCDN(c http.Client) Result {
	resp, err := GET(c, "https://api.fast.com/netflix/speedtest/v2?https=true&token=YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm&urlCount=5")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if resp.StatusCode == 403 {
		return Result{
			Status: StatusBanned,
			Info:   "IP Banned By Netflix",
		}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res netflixCdnResult
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	// u, err := url.Parse(res.Targets[0].Url)
	// if err!=nil{
	// 	return Result{Status: , Err: err}
	// }
	// ips,err:=net.LookupHost(u.Host)
	// if err!=nil{
	// 	return Result{Status: , Err: err}
	// }
	return Result{
		Status: StatusOK,
		Region: res.Targets[0].Location.Country,
	}
}

type netflixLocation struct {
	City    string
	Country string
}
type netflixCdnTarget struct {
	Name     string
	Url      string
	Location netflixLocation
}
type netflixCdnResult struct {
	Targets []netflixCdnTarget
}
