package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	// create and open file
	infoFile, err := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	warnFile, err := os.OpenFile("./warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errorFile, err := os.OpenFile("./error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// judge file if not nil
	if infoFile != nil || warnFile != nil || errorFile != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}

	// new 出对象, 赋给 *log.Logger对象
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(os.Stdout, infoFile), "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stdout, warnFile), "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errorFile), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// uuidStr := uuid.New()
	// log.SetPrefix(uuidStr.String() + " ")
	// log.SetFlags(log.Ldate | log.Lshortfile)
	// log.Print("俺是一条日志")
	// log.Printf("%s", "我是一条错误日志")
	// log.Panic("致命一击")
	// log.New(out io.Writer, prefix string, flag int)

	Info.Print("俺是一条日志")
	Warning.Println("这真的是一个log警告")
	Error.Printf("%s", "我是一条错误日志")


}
