package main

import (
	mt "MediaUnlockTest"
	"net/http"
	"sync"
	"time"
)

var (
	MUL bool
	HK  bool
	TW  bool
	JP  bool
	NA  bool
	SA  bool
)

func Check() bool {
	c := mt.AutoHttpClient
	wg := &sync.WaitGroup{}
	R = []*result{}
	if MUL {
		Multination(wg, c)
	}
	if HK {
		HongKong(wg, c)
	}
	if TW {
		Taiwan(wg, c)
	}
	if JP {
		Japan(wg, c)
	}
	if NA {
		NorthAmerica(wg, c)
	}
	if SA {

	}

	ch := make(chan struct{})
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	select {
	case <-ch:
		return false
	case <-time.After(time.Duration(Interval) * time.Second):
		return true
	}
}

type result struct {
	Type  string
	Name  string
	Value mt.Result
}

var R []*result

func excute(w *sync.WaitGroup, Name string, F func(client http.Client) mt.Result, client http.Client) {
	r := &result{Name: Name}
	R = append(R, r)
	w.Add(1)
	go func() {
		res := F(client)
		r.Value = res
		w.Done()
	}()
}

func Multination(w *sync.WaitGroup, c http.Client) {
	// R = append(R, &result{Name: "Multination", Divider: true})
	excute(w, "Dazn", mt.Dazn, c)
	excute(w, "Hotstar", mt.Hotstar, c)
	excute(w, "Disney+", mt.DisneyPlus, c)
	excute(w, "Netflix", mt.NetflixRegion, c)
	excute(w, "Netflix CDN", mt.NetflixCDN, c)
	excute(w, "Youtube", mt.YoutubeRegion, c)
	excute(w, "Youtube CDN", mt.YoutubeCDN, c)
	excute(w, "Amazon Prime Video", mt.PrimeVideo, c)
	excute(w, "TVBAnywhere+", mt.TVBAnywhere, c)
	excute(w, "iQyi", mt.IqRegion, c)
	excute(w, "Viu.com", mt.ViuCom, c)
	excute(w, "Spotify", mt.Spotify, c)
	excute(w, "Steam", mt.Steam, c)
	excute(w, "ChatGPT", mt.ChatGPT, c)
}

func HongKong(w *sync.WaitGroup, c http.Client) {
	// R = append(R, &result{Name: "Hong Kong", Divider: true})
	excute(w, "Now E", mt.NowE, c)
	excute(w, "Viu.TV", mt.ViuTV, c)
	excute(w, "MyTVSuper", mt.MyTvSuper, c)
	excute(w, "HBO GO Aisa", mt.HboGoAisa, c)
	excute(w, "BiliBili HK/Macau", mt.BilibiliHKMC, c)
}

func Taiwan(w *sync.WaitGroup, c http.Client) {
	// R = append(R, &result{Name: "Taiwan", Divider: true})
	excute(w, "KKTV", mt.KKTV, c)
	excute(w, "LiTV", mt.LiTV, c)
	excute(w, "MyVideo", mt.MyVideo, c)
	excute(w, "4GTV", mt.TW4GTV, c)
	excute(w, "LineTV", mt.LineTV, c)
	excute(w, "Hami Video", mt.HamiVideo, c)
	excute(w, "CatchPlay+", mt.Catchplay, c)
	excute(w, "Bahamut Anime", mt.BahamutAnime, c)
	excute(w, "HBO GO Aisa", mt.HboGoAisa, c)
	excute(w, "Bilibili TW", mt.BilibiliTW, c)
}

func Japan(w *sync.WaitGroup, c http.Client) {
	// R = append(R, &result{Name: "Japan", Divider: true})
	excute(w, "DMM", mt.DMM, c)
	excute(w, "DMM TV", mt.DMMTV, c)
	excute(w, "Abema", mt.Abema, c)
	excute(w, "Niconico", mt.Niconico, c)
	excute(w, "music.jp", mt.MusicJP, c)
	excute(w, "Telasa", mt.Telasa, c)
	excute(w, "Paravi", mt.Paravi, c)
	excute(w, "U-NEXT", mt.U_NEXT, c)
	excute(w, "Hulu Japan", mt.HuluJP, c)
	excute(w, "GYAO!", mt.GYAO, c)
	excute(w, "VideoMarket", mt.VideoMarket, c)
	excute(w, "FOD(Fuji TV)", mt.FOD, c)
	excute(w, "Radiko", mt.Radiko, c)
	excute(w, "Karaoke@DAM", mt.Karaoke, c)
	excute(w, "J:COM On Demand", mt.J_COM_ON_DEMAND, c)
	excute(w, "Kancolle", mt.Kancolle, c)
	excute(w, "Pretty Derby Japan", mt.PrettyDerbyJP, c)
	excute(w, "Konosuba Fantastic Days", mt.KonosubaFD, c)
	excute(w, "Princess Connect Re:Dive Japan", mt.PCRJP, c)
	excute(w, "World Flipper Japan", mt.WFJP, c)
	excute(w, "Project Sekai: Colorful Stage", mt.PJSK, c)
}

func NorthAmerica(w *sync.WaitGroup, c http.Client) {
	// R = append(R, &result{Name: "North America", Divider: true})
	excute(w, "FOX", mt.Fox, c)
	excute(w, "Hulu", mt.Hulu, c)
	excute(w, "ESPN+", mt.ESPNPlus, c)
	excute(w, "Epix", mt.Epix, c)
	excute(w, "Starz", mt.Starz, c)
	excute(w, "Philo", mt.Philo, c)
	excute(w, "FXNOW", mt.FXNOW, c)
	excute(w, "TLC GO", mt.TlcGo, c)
	excute(w, "HBO Max", mt.HBOMax, c)
	excute(w, "Shudder", mt.Shudder, c)
	excute(w, "BritBox", mt.BritBox, c)
	excute(w, "CW TV", mt.CW_TV, c)
	excute(w, "NBA TV", mt.NBA_TV, c)
	excute(w, "Fubo TV", mt.FuboTV, c)
	excute(w, "Tubi TV", mt.TubiTV, c)
	excute(w, "Sling TV", mt.SlingTV, c)
	excute(w, "Pluto TV", mt.PlutoTV, c)
	excute(w, "Acorn TV", mt.AcornTV, c)
	excute(w, "SHOWTIME", mt.SHOWTIME, c)
	excute(w, "encoreTVB", mt.EncoreTVB, c)
	excute(w, "Funimation", mt.Funimation, c)
	excute(w, "Discovery+", mt.DiscoveryPlus, c)
	excute(w, "Paramount+", mt.ParamountPlus, c)
	excute(w, "Peacock TV", mt.PeacockTV, c)
	excute(w, "Popcornflix", mt.Popcornflix, c)
	excute(w, "Crunchyroll", mt.Crunchyroll, c)
	excute(w, "Direct Stream", mt.DirectvStream, c)
	// R = append(R, &result{Name: "CA", Divider: true})
	excute(w, "CBC Gem", mt.CBCGem, c)
	excute(w, "Crave", mt.Crave, c)
}
