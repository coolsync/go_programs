package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 遍历 命令行参数， 获取 url地址
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		// 发送请求， 获取响应包
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			// 读取response pkg, get data
			// bytes, err := ioutil.ReadAll(resp.Body)
			_, err = io.Copy(os.Stdout, resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "io.Copy: %v\n", err)
				os.Exit(1)
			}
			// print data info
			// fmt.Printf("data：\n%s\n", bytes)
		} else {
			fmt.Fprintf(os.Stderr, "status code %d\n", resp.StatusCode)
		}

		// fmt.Printf("\nstatus code: %d\n", resp.StatusCode)
	}
}
