package main

import (
	"os"
	"strconv"
)

func main() {
	var rmdirs []func()

	for _, d := range tmpDir() {
		d := d
		os.MkdirAll(d, 0755)

		rmdirs = append(rmdirs, func() {
			os.RemoveAll(d)
		})
	}

	// do more work

	for _, rmdir := range rmdirs {
		rmdir()
	}

}

// create most temp dir
func tmpDir() []string {
	var dirs []string
	for i := 0; i < 10; i++ {
		dirname := "./tmp-" + strconv.Itoa(i+1)
		dirs = append(dirs, dirname)
	}
	return dirs
}
