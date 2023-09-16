package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Service() {
	args := []string{}
	for _, a := range os.Args[1:] {
		if !strings.Contains(a, "-service") {
			args = append(args, a)
		}
	}
	data := []byte(`[Unit]
Description=unlock-monitor
After=network.target
[Service]
Type=simple
LimitCPU=infinity
LimitFSIZE=infinity
LimitDATA=infinity
LimitSTACK=infinity
LimitCORE=infinity
LimitRSS=infinity
LimitNOFILE=infinity
LimitAS=infinity
LimitNPROC=infinity
LimitMEMLOCK=infinity
LimitLOCKS=infinity
LimitSIGPENDING=infinity
LimitMSGQUEUE=infinity
LimitRTPRIO=infinity
LimitRTTIME=infinity
ExecStart=/usr/bin/unlock-monitor ` + strings.Join(args, " ") + `
Restart=always
RestartSec=5
[Install]
WantedBy=multi-user.target`)
	if err := ioutil.WriteFile("/etc/systemd/system/unlock-monitor.service", data, 0644); err != nil {
		log.Fatal("[ERR] 写入systemd守护service时出错:", err)
	}
	log.Println("[OK] systemd守护service成功")
	if err := RunCmd("systemctl", "daemon-reload"); err != nil {
		log.Fatal("[ERR] 重载systemctl时出错:", err)
	}
	if err := RunCmd("systemctl", "restart", "unlock-monitor"); err != nil {
		log.Fatal("[ERR] 启动tun服务时出错:", err)
	}
	if err := RunCmd("systemctl", "enable", "unlock-monitor"); err != nil {
		log.Fatal("[ERR] 设置tun服务开机自启时出错:", err)
	}
	log.Println("[OK] 初始化服务成功")
}
