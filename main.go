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

var bar *progressbar.ProgressBar
var wg = &sync.WaitGroup{}

func excute(Name string, F func(c http.Client) Result, C http.Client) {
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
			return FontYellow + "ERR: (" + r.Err.Error() + ")" + FontSuffix
		} else if r.Info != "" {
			return FontRed + " NO (" + r.Info + ")" + FontSuffix
		} else {
			return FontRed + " NO" + FontSuffix
		}
	}
	return
}

func main() {
	// log.Println(Abema(Ipv4HttpClient))
	// return
	bar = progressbar.Default(48, "testing ...")

	excute("Dazn", Dazn, Ipv4HttpClient)
	excute("Hotstar", Hotstar, Ipv4HttpClient)
	excute("Disney+", DisneyPlus, Ipv4HttpClient)
	excute("Netflix", NetflixRegion, Ipv4HttpClient)
	excute("Netflix CDN", NetflixCDN, Ipv4HttpClient)
	excute("Youtube", YoutubeRegion, Ipv4HttpClient)
	excute("Youtube CDN", YoutubeCDN, Ipv4HttpClient)
	excute("Prime Video", PrimeVideo, Ipv4HttpClient)
	excute("TVBAnywhere+", TVBAnywhere, Ipv4HttpClient)
	excute("iQyi Region", IqRegion, Ipv4HttpClient)
	excute("Viu.com", ViuCom, Ipv4HttpClient)
	excute("Spotify", Spotify, Ipv4HttpClient)
	excute("Steam", Steam, Ipv4HttpClient)

	excute("Viu.TV", ViuTV, Ipv4HttpClient)
	excute("Now E", NowE, Ipv4HttpClient)
	excute("MyTVSuper", MyTvSuper, Ipv4HttpClient)
	excute("HBO GO Aisa", HboGoAisa, Ipv4HttpClient)
	excute("BiliBili Hongkong/Macau/Taiwan", BilibiliHKMCTW, Ipv4HttpClient)

	excute("KKTV", KKTV, Ipv4HttpClient)
	excute("LiTV", LiTV, Ipv4HttpClient)
	excute("MyVideo", MyVideo, Ipv4HttpClient)
	excute("TW4GTV", TW4GTV, Ipv4HttpClient)
	excute("LineTV", LineTV, Ipv4HttpClient)
	excute("HamiVideo", HamiVideo, Ipv4HttpClient)
	excute("Catchplay+", Catchplay, Ipv4HttpClient)
	excute("Bahamu Anime", BahamuAnime, Ipv4HttpClient)
	excute("HBO GO Aisa", HboGoAisa, Ipv4HttpClient)
	excute("Bilibili Taiwan Only", BilibiliTW, Ipv4HttpClient)

	excute("DMM", DMM, Ipv4HttpClient)
	excute("Abema", Abema, Ipv4HttpClient)
	excute("Niconico", Niconico, Ipv4HttpClient)
	excute("music.jp", MusicJP, Ipv4HttpClient)
	excute("Telasa", Telasa, Ipv4HttpClient)
	excute("Paravi", Paravi, Ipv4HttpClient)
	excute("U-NEXT", U_NEXT, Ipv4HttpClient)
	excute("Hulu Japan", HuluJP, Ipv4HttpClient)
	excute("GYAO!", GYAO, Ipv4HttpClient)
	excute("VideoMarket", VideoMarket, Ipv4HttpClient)
	excute("FOD(Fuji TV)", FOD, Ipv4HttpClient)
	excute("Radiko", Radiko, Ipv4HttpClient)
	excute("Karaoke@DAM", Karaoke, Ipv4HttpClient)
	excute("J:COM On Demand", J_COM_ON_DEMAND, Ipv4HttpClient)
	excute("Kancolle", Kancolle, Ipv4HttpClient)
	excute("Pretty Derby Japan", PrettyDerbyJP, Ipv4HttpClient)
	excute("Konosuba Fantastic Days", KonosubaFD, Ipv4HttpClient)
	excute("Princess Connect Re:Dive Japan", PCRJP, Ipv4HttpClient)
	excute("World Flipper Japan", WFJP, Ipv4HttpClient)
	excute("Project Sekai: Colorful Stage", PJSK, Ipv4HttpClient)

	wg.Wait()
	bar.Describe("Finished")
	bar.Clear()
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
