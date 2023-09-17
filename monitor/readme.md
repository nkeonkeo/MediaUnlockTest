## unlock-monitor

对接 prometheus

效果:

![]()

安装: 

```bash
bash <(curl -Ls unlock.moe/monitor) -service
```

使用:

```
Usage of unlock-monitor:
  -interval int
        check interval (s) (default 60)
  -service
        setup systemd service
  -hk
        Hong Kong
  -jp
        Japan
  -listen string
        listen address (default ":9101")
  -mul
        Multination (default true)
  -na
        North America
  -sa
        South America
  -tw
        Taiwan
  -u    check update
  -v    show version
```