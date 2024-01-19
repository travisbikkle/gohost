package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)
import "net/http"

type IpResponse struct {
	Address string   `json:"address"`
	Ttl     float64  `json:"ttl"`
	Reverse []string `json:"reverse"`
}

var (
	server string
	urls   []string
)

func main() {
	bindFlags()
	host := "/etc/hosts"
	if runtime.GOOS == "windows" {
		host = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	}
	fmt.Printf("copy following lines to %s\n", host)
	for _, domain := range urls {
		do(domain)
	}
}

func bindFlags() {
	flag.StringVar(&server, "s", "https://dnschecker.fabdev.eu.org", "server to query")
	flag.Usage = func() {
		fmt.Println("Usage: $0 [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()
	urls = flag.Args()
	if len(urls) == 0 {
		urls = []string{"github.com", "github.global.ssl.fastly.net", "assets-cdn.github.com", "codeload.github.com", "github.io"}
	}
}

func do(domain string) {
	ip := GetIp(domain)
	//WriteToHosts(ip, domain)
	fmt.Printf("%-18s%s\n", ip, domain)
}

func GetIp(domain string) (ip string) {
	ip = ""
	url := fmt.Sprintf("%s/A/%s", server, domain)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var ipResponse []IpResponse
		err = json.Unmarshal(bodyBytes, &ipResponse)
		if err != nil {
			log.Fatal(err)
		}
		ip = ipResponse[0].Address
	}
	return
}
