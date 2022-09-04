package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/shirou/gopsutil/host"
)

var link = "https://webhook.site/3da58b43-356c-4bff-b48e-df0ca1a0417d"

func main() {
	hostStat, _ := host.Info()

	fmt.Printf("%+v\n", hostStat)
	data := url.Values{
		"host": {hostStat.HostID},
	}

	_, err := http.PostForm(link, data)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	} else {
		log.Printf("Done :)")
	}
}
