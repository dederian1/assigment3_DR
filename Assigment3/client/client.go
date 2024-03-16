package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for {
		// Kirim permintaan update ke server
		resp, err := http.Get("http://localhost:8080/update")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Baca body respons
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Tampilkan respons
		fmt.Println(string(body))

		// Tutup body respons
		resp.Body.Close()

		// Tunggu selama 15 detik
		time.Sleep(15 * time.Second)
	}
}
