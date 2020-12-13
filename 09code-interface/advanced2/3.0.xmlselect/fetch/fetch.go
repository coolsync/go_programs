package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		err := outlink(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func outlink(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http get %v", err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return fmt.Errorf("io.Copy err: %v", err)
	}

	return nil
}
