package main

import (
	"flag"
	"fmt"
	"go-programs/08code-function/05-anonymous-function/links"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// 解析 命令行 传入的 params
// 将 param 传入 crawl, 用于获取一个页面 all links
// 创建 wg, 用于并发处理 download
// download, 用于获取页面单个链接中的内容
// 获取响应包内容,
// 创建目录和对应的内容存储的html
// io.Copy(), 传入内容

var base = flag.String("base", "https://golangbyexample.com", "请输入 query url")
var wg sync.WaitGroup

func download(base, url string) {
	defer wg.Done()

	// url 与 base 是否一致
	if url == base {
		return
	}
	// if !strings.HasPrefix(base, url) {
	// 	return
	// }

	// 获取 url相应包
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// 创建 mkdir and file
	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Println(err)
	}

	// 创建文件
	filename := dir + "/index.html"

	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	// 将应答包内容写入文件
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Println(err)
	}
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	fmt.Println(list)
	if err != nil {
		log.Println(err)
	}

	return list
}

func main() {
	flag.Parse()

	for _, url := range crawl(*base) {
		wg.Add(1)
		go download(*base, url)
	}

	done := make(chan struct{})

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	<-done
}
