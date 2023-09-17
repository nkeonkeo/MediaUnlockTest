package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	pb "github.com/schollz/progressbar/v3"
)

var (
	Version   = "1.0"
	buildTime string
)

func checkUpdate() {
	resp, err := http.Get("https://unlock.moe/monitor/latest/version")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	version := string(b)
	if version == Version {
		fmt.Println("已经是最新版本")
		return
	}
	fmt.Println("检测到新版本", version)
	OS, ARCH := runtime.GOOS, runtime.GOARCH
	fmt.Println("OS:", OS)
	fmt.Println("ARCH:", ARCH)
	out, err := os.Create("/usr/bin/unlock-monitor_new")
	if err != nil {
		log.Fatal("[ERR] 创建文件出错:", err)
		return
	}
	defer out.Close()
	log.Println("下载unlock-monitor中 ...")
	url := "https://unlock.moe/monitor/latest/unlock-monitor_" + runtime.GOOS + "_" + runtime.GOARCH
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal("[ERR] 下载unlock-monitor时出错:", err)
	}
	defer resp.Body.Close()
	downloader := &Downloader{
		Reader: resp.Body,
		Total:  uint64(resp.ContentLength),
		Pb:     pb.DefaultBytes(resp.ContentLength, "下载进度"),
	}
	if _, err := io.Copy(out, downloader); err != nil {
		log.Fatal("[ERR] 下载unlock-monitor时出错:", err)
	}
	if os.Chmod("/usr/bin/unlock-monitor_new", 0777) != nil {
		log.Fatal("[ERR] 更改unlock-monitor后端权限出错:", err)
	}
	if _, err := os.Stat("/usr/bin/unlock-monitor"); err == nil {
		if os.Remove("/usr/bin/unlock-monitor") != nil {
			log.Fatal("[ERR] 删除unlock-monitor旧版本时出错:", err.Error())
		}
	}
	if os.Rename("/usr/bin/unlock-monitor_new", "/usr/bin/unlock-monitor") != nil {
		log.Fatal("[ERR] 更新unlock-monitor后端时出错:", err)
	}
	log.Println("[OK] unlock-monitor后端更新成功")
}

type Downloader struct {
	io.Reader
	Total   uint64
	Current uint64
	Pb      *pb.ProgressBar
	done    bool
}

func (d *Downloader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	d.Current += uint64(n)
	if d.done {
		return
	}
	d.Pb.Add(n)
	// fmt.Printf("\r正在下载，进度：%.2f%% [%s/%s]", float64(d.Current*10000/d.Total)/100, humanize.Bytes(d.Current), humanize.Bytes(d.Total))
	if d.Current == d.Total {
		d.done = true
		d.Pb.Describe("unlock-monitor下载完成")
		d.Pb.Finish()
	}
	return
}
