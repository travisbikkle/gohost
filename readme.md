# gohost
gohost resolves those domain names affected by DNS cache poisoning (also known as DNS spoofing) and returns true IPs.
If a domain name points to multiple IPs, only the first IP will be returned. 

It's very useful when accessing github related sites and clone/push code from/to github, in some special places.

## Build from source
```bash
## build locally
go build -o gohost
## cross compile on Linux for Windows or MacOS, might get false positive from windows defender --> "Trojan:Win32xxx"
GOOS=windows GOARCH=amd64 go build -o gohost_amd64.exe
GOOS=darwin GOARCH=amd64 go build -o gohost_darwin_amd64
```

## Or download from releases and run
```bash
D:\Downloads>.\gohost.exe -s
flag needs an argument: -s
Usage: $0 [options]
Options:
  -s string
        server to query (default "https://dnschecker.fabdev.eu.org")

D:\Downloads>.\gohost.exe
copy fowllowing lines to C:\Windows\System32\drivers\etc\hosts
140.82.121.3      github.com
151.101.1.194     github.global.ssl.fastly.net
185.199.108.153   assets-cdn.github.com
140.82.121.10     codeload.github.com
185.199.110.153   github.io

D:\Downloads>.\gohost.exe www.google.com www.xxxx.com
copy fowllowing lines to C:\Windows\System32\drivers\etc\hosts
142.250.185.132   www.google.com
103.235.46.40     www.xxxx.com
```

## Execute additional command on windows
```bash
ipconfig /flushdns
```

## You might also want to set
```bash
git config --global http.sslVerify "false"
```
