package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/schollz/progressbar/v3"
)

type result struct {
	Name  string
	Value Result
}

var R = []*result{}

var bar = progressbar.Default(-1, "testing ...")

func excute(Name string, F func(c http.Client) Result, C http.Client, wg *sync.WaitGroup) {
	r := &result{Name: Name}
	R = append(R, r)
	wg.Add(1)
	go func() {
		r.Value = F(C)
		bar.Describe(Name + " " + ShowResult(r.Value))
		bar.Add(1)
		wg.Done()
	}()
}

var (
	FontRed     = "\033[31m"
	FontGreen   = "\033[32m"
	FontYellow  = "\033[33m"
	FontBlue    = "\033[34m"
	FontPurple  = "\033[35m"
	FontSkyBlue = "\033[36m"
	FontWhite   = "\033[37m"
	FontSuffix  = "\033[0m"
)

func ShowResult(r Result) (s string) {
	if r.Success {
		s = FontGreen + "YES"
		if r.Region != "" {
			s += " (Region: " + r.Region + ")"
		}
		s += FontSuffix
	} else {
		if r.Err != nil {
			return FontYellow + "ERR: " + r.Err.Error() + ")" + FontSuffix
		} else if r.Info != "" {
			return FontRed + "NO (" + r.Info + ")" + FontSuffix
		} else {
			return FontRed + "NO" + FontSuffix
		}
	}
	return
}

func main() {

	wg := &sync.WaitGroup{}
	excute("Dazn", Dazn, Ipv4HttpClient, wg)
	excute("Hotstar", Hotstar, Ipv4HttpClient, wg)
	excute("DisneyPlus", DisneyPlus, Ipv4HttpClient, wg)
	excute("Netflix", NetflixRegion, Ipv4HttpClient, wg)
	excute("NetflixCDN", NetflixCDN, Ipv4HttpClient, wg)
	excute("Youtube", YoutubeRegion, Ipv4HttpClient, wg)
	excute("YoutubeCDN", YoutubeCDN, Ipv4HttpClient, wg)
	excute("PrimeVideo", PrimeVideo, Ipv4HttpClient, wg)
	excute("TVBAnywhere", TVBAnywhere, Ipv4HttpClient, wg)
	excute("IqRegion", IqRegion, Ipv4HttpClient, wg)
	excute("ViuCom", ViuCom, Ipv4HttpClient, wg)
	excute("Spotify", Spotify, Ipv4HttpClient, wg)

	excute("Steam", Steam, Ipv4HttpClient, wg)
	excute("ViuTV", ViuTV, Ipv4HttpClient, wg)
	excute("NowE", NowE, Ipv4HttpClient, wg)
	excute("MyTvSuper", MyTvSuper, Ipv4HttpClient, wg)
	excute("HboGoAisa", HboGoAisa, Ipv4HttpClient, wg)
	excute("BilibiliHKMCTW", BilibiliHKMCTW, Ipv4HttpClient, wg)

	excute("DMM", DMM, Ipv4HttpClient, wg)
	excute("Abema", Abema, Ipv4HttpClient, wg)
	excute("Niconico", Niconico, Ipv4HttpClient, wg)
	excute("MusicJP", MusicJP, Ipv4HttpClient, wg)
	excute("Telasa", Telasa, Ipv4HttpClient, wg)
	excute("Paravi", Paravi, Ipv4HttpClient, wg)
	excute("U_NEXT", U_NEXT, Ipv4HttpClient, wg)
	excute("HuluJP", HuluJP, Ipv4HttpClient, wg)
	excute("GYAO", GYAO, Ipv4HttpClient, wg)
	excute("VideoMarket", VideoMarket, Ipv4HttpClient, wg)
	excute("FOD", FOD, Ipv4HttpClient, wg)
	excute("Radiko", Radiko, Ipv4HttpClient, wg)
	excute("Karaoke", Karaoke, Ipv4HttpClient, wg)
	excute("J_COM_ON_DEMAND", J_COM_ON_DEMAND, Ipv4HttpClient, wg)
	excute("Kancolle", Kancolle, Ipv4HttpClient, wg)
	excute("PrettyDerbyJP", PrettyDerbyJP, Ipv4HttpClient, wg)
	excute("KonosubaFD", KonosubaFD, Ipv4HttpClient, wg)
	excute("PCRJP", PCRJP, Ipv4HttpClient, wg)
	excute("WFJP", WFJP, Ipv4HttpClient, wg)
	excute("PJSK", PJSK, Ipv4HttpClient, wg)

	excute("KKTV", KKTV, Ipv4HttpClient, wg)
	excute("LiTV", LiTV, Ipv4HttpClient, wg)
	excute("MyVideo", MyVideo, Ipv4HttpClient, wg)
	excute("TW4GTV", TW4GTV, Ipv4HttpClient, wg)
	excute("LineTV", LineTV, Ipv4HttpClient, wg)
	excute("HamiVideo", HamiVideo, Ipv4HttpClient, wg)
	excute("Catchplay", Catchplay, Ipv4HttpClient, wg)
	excute("BahamuAnime", BahamuAnime, Ipv4HttpClient, wg)
	excute("HboGoAisa", HboGoAisa, Ipv4HttpClient, wg)
	excute("BilibiliTW", BilibiliTW, Ipv4HttpClient, wg)

	wg.Wait()
	bar.Describe("Finished")
	bar.Finish()

	NameLength := 0
	for _, r := range R {
		if len(r.Name) > NameLength {
			NameLength = len(r.Name)
		}
	}
	for _, r := range R {
		fmt.Printf("%-"+strconv.Itoa(NameLength)+"s %s\n", r.Name, ShowResult(r.Value))
	}
}
