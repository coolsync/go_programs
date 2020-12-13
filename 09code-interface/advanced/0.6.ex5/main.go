package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
/* 
type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
} 
*/

func main() {
	lr := LimitReader(strings.NewReader("12345"), 2) // 限制读取输入字符串前2位

	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "readall err %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", b)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

type LimitedReader struct {
	underlying  io.Reader
	remainBytes int64
}

func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	if lr.remainBytes <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.remainBytes {
		p = p[:lr.remainBytes]
	}

	n, err = lr.underlying.Read(p)

	lr.remainBytes -= int64(n)

	return n, err
}

