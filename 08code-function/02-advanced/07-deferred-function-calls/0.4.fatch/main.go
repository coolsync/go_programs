package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// os.Args, 获取url
// get response pkg
// get url 路径 最后 /
// io.Copy 将 url 的结果写入文件
// 将 err, n, filname 返回

func fetch(url string) (string, int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)	// 获取 / 以及后面的文件名
	
	if !strings.HasSuffix(local, "/") {
		local = "/"
	}

	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err := io.Copy(f, resp.Body)

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}

func main(){
	for _, url := range os.Args[1:] {
		filename, size, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err: %v\n", err)
		}

		fmt.Fprintf(os.Stdout, "file name: %s\tfile size:%d\n",filename, size)
	}
}