package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

// str format 解析成 time format
func length(s string) time.Duration { // type Duration int64
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

// 按 tab format 打印输出数据
func printTracks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\n"
	tbw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight) // 创建tab输出格式, 并按右对齐
	fmt.Fprintf(tbw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tbw, format, "------", "------", "------", "------", "------")

	for _, t := range tracks {
		fmt.Fprintf(tbw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tbw.Flush()
}

// 创建多层排序功能， Title, Year, 运行时间Length依次排序
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// 实现 sort.Interface
func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func main() {
	fmt.Println("customSort sort before:")
	printTracks(tracks)

	// 转换 tracks 成customSort 实现sort.Interface
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title // title 一致， 不进行比较
		}
		if x.Year != y.Year {
			return x.Year < y.Year // year 一致 不比较
		}
		if x.Length != y.Length { // 运行时间Length 一致 不比较
			return x.Length < y.Length
		}

		return false
	}})

	fmt.Println("\ncustomSort sort after:")
	printTracks(tracks)
}
