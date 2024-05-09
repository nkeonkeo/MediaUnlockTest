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
	EU  bool
)

type TEST struct {
	Client  http.Client
	Results []*result
	Wg      *sync.WaitGroup
}

func NewTest() *TEST {
	t := &TEST{
		Client:  mt.NewAutoHttpClient(),
		Results: make([]*result, 0),
		Wg:      &sync.WaitGroup{},
	}
	return t
}

func (T *TEST) Check() bool {
	if MUL {
		T.Multination()
	}
	if HK {
		T.HongKong()
	}
	if TW {
		T.Taiwan()
	}
	if JP {
		T.Japan()
	}
	if NA {
		T.NorthAmerica()
	}
	if SA {
        T.SouthAmerica()
	}
	if EU {
		T.Europe()
	}

	ch := make(chan struct{})
	go func() {
		defer close(ch)
		T.Wg.Wait()
	}()
	select {
	case <-ch:
		return false
	case <-time.After(30 * time.Second):
		return true
	}
}

type result struct {
	Type  string
	Name  string
	Value mt.Result
}

func (T *TEST) excute(Name string, F func(client http.Client) mt.Result) {
	r := &result{Name: Name}
	T.Results = append(T.Results, r)
	T.Wg.Add(1)
	go func() {
		res := F(T.Client)
		r.Value = res
		T.Wg.Done()
	}()
}

func (T *TEST) Multination() {
	// R = append(R, &result{Name: "Multination", Divider: true})
	T.excute("Dazn", mt.Dazn)
	T.excute("Hotstar", mt.Hotstar)
	T.excute("Disney+", mt.DisneyPlus)
	T.excute("Netflix", mt.NetflixRegion)
	T.excute("Netflix CDN", mt.NetflixCDN)
	T.excute("Youtube", mt.YoutubeRegion)
	T.excute("Youtube CDN", mt.YoutubeCDN)
	T.excute("Amazon Prime Video", mt.PrimeVideo)
	T.excute("TVBAnywhere+", mt.TVBAnywhere)
	T.excute("iQyi", mt.IqRegion)
	T.excute("Viu.com", mt.ViuCom)
	T.excute("Spotify", mt.Spotify)
	T.excute("Steam", mt.Steam)
	T.excute("ChatGPT", mt.ChatGPT)
	T.excute("Wikipedia", mt.WikipediaEditable)
	T.excute("Reddit", mt.Reddit)
}

func (T *TEST) HongKong() {
	// R = append(R, &result{Name: "Hong Kong", Divider: true})
	T.excute("Now E", mt.NowE)
	T.excute("Viu.TV", mt.ViuTV)
	T.excute("MyTVSuper", mt.MyTvSuper)
	T.excute("HBO GO Aisa", mt.HboGoAisa)
	T.excute("BiliBili HK/Macau", mt.BilibiliHKMC)
}

func (T *TEST) Taiwan() {
	// R = append(R, &result{Name: "Taiwan", Divider: true})
	T.excute("KKTV", mt.KKTV)
	T.excute("LiTV", mt.LiTV)
	T.excute("MyVideo", mt.MyVideo)
	T.excute("4GTV", mt.TW4GTV)
	T.excute("LineTV", mt.LineTV)
	T.excute("Hami Video", mt.HamiVideo)
	T.excute("CatchPlay+", mt.Catchplay)
	T.excute("Bahamut Anime", mt.BahamutAnime)
	T.excute("HBO GO Aisa", mt.HboGoAisa)
	T.excute("Bilibili TW", mt.BilibiliTW)
}

func (T *TEST) Japan() {
	// R = append(R, &result{Name: "Japan", Divider: true})
	T.excute("DMM", mt.DMM)
	T.excute("DMM TV", mt.DMMTV)
	T.excute("Abema", mt.Abema)
	T.excute("Niconico", mt.Niconico)
	T.excute("music.jp", mt.MusicJP)
	T.excute("Telasa", mt.Telasa)
	T.excute("Paravi", mt.Paravi)
	T.excute("U-NEXT", mt.U_NEXT)
	T.excute("Hulu Japan", mt.HuluJP)
	T.excute("GYAO!", mt.GYAO)
	T.excute("VideoMarket", mt.VideoMarket)
	T.excute("FOD(Fuji TV)", mt.FOD)
	T.excute("Radiko", mt.Radiko)
	T.excute("Karaoke@DAM", mt.Karaoke)
	T.excute("J:COM On Demand", mt.J_COM_ON_DEMAND)
	T.excute("Kancolle", mt.Kancolle)
	T.excute("Pretty Derby Japan", mt.PrettyDerbyJP)
	T.excute("Konosuba Fantastic Days", mt.KonosubaFD)
	T.excute("Princess Connect Re:Dive Japan", mt.PCRJP)
	T.excute("World Flipper Japan", mt.WFJP)
	T.excute("Project Sekai: Colorful Stage", mt.PJSK)
}

func (T *TEST) NorthAmerica() {
	// R = append(R, &result{Name: "North America", Divider: true})
	T.excute("FOX", mt.Fox)
	T.excute("Hulu", mt.Hulu)
	T.excute("ESPN+", mt.ESPNPlus)
	T.excute("Epix", mt.Epix)
	T.excute("Starz", mt.Starz)
	T.excute("Philo", mt.Philo)
	T.excute("FXNOW", mt.FXNOW)
	T.excute("TLC GO", mt.TlcGo)
	T.excute("HBO Max", mt.HBOMax)
	T.excute("Shudder", mt.Shudder)
	T.excute("BritBox", mt.BritBox)
	T.excute("CW TV", mt.CW_TV)
	T.excute("NBA TV", mt.NBA_TV)
	T.excute("Fubo TV", mt.FuboTV)
	T.excute("Tubi TV", mt.TubiTV)
	T.excute("Sling TV", mt.SlingTV)
	T.excute("Pluto TV", mt.PlutoTV)
	T.excute("Acorn TV", mt.AcornTV)
	T.excute("SHOWTIME", mt.SHOWTIME)
	T.excute("encoreTVB", mt.EncoreTVB)
	T.excute("Funimation", mt.Funimation)
	T.excute("Discovery+", mt.DiscoveryPlus)
	T.excute("Paramount+", mt.ParamountPlus)
	T.excute("Peacock TV", mt.PeacockTV)
	T.excute("Popcornflix", mt.Popcornflix)
	T.excute("Crunchyroll", mt.Crunchyroll)
	T.excute("Direct Stream", mt.DirectvStream)
	// R = append(R, &result{Name: "CA", Divider: true})
	T.excute("CBC Gem", mt.CBCGem)
	T.excute("Crave", mt.Crave)
}

func (T *TEST) SouthAmerica() {
    //R = append(R, &result{Name: "South America", Divider: true})
    T.excute("Star Plus", mt.StarPlus)
    T.excute("DirecTV GO", mt.DirecTVGO)
}

func (T *TEST) Europe() {
    //R = append(R, &result{Name: "Europe", Divider: true})
    T.excute("BBC iPlayer", mt.BBCiPlayer)
    T.excute("Rakuten TV", mt.RakutenTV)
    //excute("Sky Show Time", m.SkyShowTime, c)
}
