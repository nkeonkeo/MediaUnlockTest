## unlock-monitor

使用 Grafana 对接 Prometheus，搭建流媒体监控页面。

### 效果:

![](https://raw.githubusercontent.com/HsukqiLee/MediaUnlockTest/main/monitor/image.png)

### 安装: 

```bash
bash <(curl -Ls unlock.icmp.ing/monitor.sh) -service
```

### 使用:

```
Usage of unlock-monitor:
  -listen string
        listen address (default ":9101")
  -interval int
        check interval (s) (default 60)
  -service
        setup systemd service
  -mul
        Multination (default true)
  -hk
        Hong Kong
  -tw
        Taiwan
  -jp
        Japan
  -na
        North America
  -sa
        South America
  -eu
        Europe
  -u    check update
  -v    show version
```

### Prometheus: 

添加 Job:

```yaml
- job_name: checkmedia
    scrape_interval: 30s
    static_configs:
      - targets:
        - <your ip/domain>:9101
        - ...
```

### Grafana

Value mappings

|Status|Display Text|
|---|---|
|1|YES|
|2|Restricted|
|3|NO|
|4|BANNED|
|5|FAILED|
|6|UNEXPECTED|
|-1|NET ERR|
|-2|ERROR|

经测试在网络异常（触发原因不明，因为本来应该返回`-1`）时`status`字段可能为`0`，可显示为`Unknown`或`TIMEOUT`等。
