package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http_prefix := "http://"
	https_prefix := "https://"
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, http_prefix) || strings.HasPrefix(url, https_prefix)) {
			url = "http://" + url
		}
		rsp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, rsp.Body)
		rsp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			os.Exit(1)
		}
		fmt.Println("Status code =", rsp.Status)
	}
}
