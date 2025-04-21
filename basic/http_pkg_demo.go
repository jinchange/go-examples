package basic

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HttpClientDemo() {
	resp, err := http.Get("https://httbafbxin.org/isp")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	println(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
