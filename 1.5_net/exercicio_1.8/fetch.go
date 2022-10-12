package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		urlFormatted := formatURL(url)
		fmt.Println(urlFormatted)

		resp, err := http.Get(urlFormatted)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, resp.Status)
		_, err = io.CopyBuffer(os.Stdout, resp.Body, make([]byte, 1))
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func formatURL(uri string) string {

	if !strings.HasPrefix(uri, "http://") {
		uri = "http://" + uri
	}
	return uri

}
