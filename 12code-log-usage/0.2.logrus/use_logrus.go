package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func initLog() {
	// 设置格式为 json格式
	logrus.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})
}

func main() {
	initLog()

	logrus.WithFields(logrus.Fields{
		"age":  12,
		"name": "xiaohei",
		"sex":  1,
	}).Info("小黑来了")

	logrus.WithFields(logrus.Fields{
		"age":  13,
		"name": "xiaohong",
		"sex":  0,
	}).Error("小红来了")

	// logrus.WithFields(logrus.Fields{
	// 	"age":  14,
	// 	"name": "xiaofang",
	// 	"sex":  0,
	// }).Fatal("小芳来了")

	f, err := os.OpenFile("./1.log", 64|1, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// 指定日志输出的地方
	logrus.SetOutput(f)
	logrus.SetLevel(logrus.InfoLevel)

}

func funcFlag2Int() {
	const (
		// Invented values to support what package os expects.
		O_RDONLY   = 0x00000
		O_WRONLY   = 0x00001
		O_RDWR     = 0x00002
		O_CREAT    = 0x00040
		O_EXCL     = 0x00080
		O_NOCTTY   = 0x00100
		O_TRUNC    = 0x00200
		O_NONBLOCK = 0x00800
		O_APPEND   = 0x00400
		O_SYNC     = 0x01000
		O_ASYNC    = 0x02000
		O_CLOEXEC  = 0x80000
	)

	fmt.Printf("O_RDONLY: %d\n", O_RDONLY)     // O_RDONLY: 0
	fmt.Printf("O_WRONLY: %d\n", O_WRONLY)     // O_WRONLY: 1
	fmt.Printf("O_RDWR: %d\n", O_RDWR)         // O_RDWR: 2
	fmt.Printf("O_CREAT: %d\n", O_CREAT)       // O_CREAT: 64
	fmt.Printf("O_EXCL: %d\n", O_EXCL)         // O_EXCL: 128
	fmt.Printf("O_NOCTTY: %d\n", O_NOCTTY)     // O_NOCTTY: 256
	fmt.Printf("O_TRUNC: %d\n", O_TRUNC)       // O_TRUNC: 512
	fmt.Printf("O_NONBLOCK: %d\n", O_NONBLOCK) // O_NONBLOCK: 2048
	fmt.Printf("O_APPEND: %d\n", O_APPEND)     // O_APPEND: 1024
	fmt.Printf("O_SYNC: %d\n", O_SYNC)         // O_SYNC: 4096
	fmt.Printf("O_ASYNC: %d\n", O_ASYNC)       // O_ASYNC: 8192
	fmt.Printf("O_CLOEXEC: %d\n", O_CLOEXEC)   // O_CLOEXEC: 524288
}
