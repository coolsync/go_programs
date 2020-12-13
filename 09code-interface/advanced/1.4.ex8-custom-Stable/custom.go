package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// define 专辑信息的 struct
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// 创建 Track 集合
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

// 转换 time format
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// 创建 table format output
func printTracks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\n"
	tbw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 0, '\t', 0)

	// 实现io.Writer Write method
	fmt.Fprintf(tbw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tbw, format, "------", "------", "------", "------", "------")

	for _, v := range tracks {
		fmt.Fprintf(tbw, format, v.Title, v.Artist, v.Album, v.Year, v.Length)
	}

	tbw.Flush()
}

// 创建 custom sort, 用于完成 sort.Stable形式的配对
// 提取各个键的排序, 非整体排序
type custom struct {
	t    []*Track
	less func(x, y *Track) bool
	swap func(x, y *Track)
}

// 实现 sort.Interface
func (c custom) Len() int {
	return len(c.t)
}

func (c custom) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c custom) Swap(i, j int) {
	c.swap(c.t[i], c.t[j])
}

// 创建 click func, 指定主键及其后续键
func click(s string) {
	switch s {
	case "title":
		sort.Stable(custom{tracks, // 传入struct, 使用 sort.Stable 排序
			func(x, y *Track) bool {
				return x.Title < y.Title
			},
			func(x, y *Track) {
				x.Title, y.Title = y.Title, x.Title
			},
		})
	case "artist":
		sort.Stable(custom{tracks, // 传入struct, 使用 sort.Stable 排序
			func(x, y *Track) bool {
				return x.Artist < y.Artist
			},
			func(x, y *Track) {
				fmt.Printf("\n%s, %s\n\n", x.Artist, y.Artist)
				x.Artist, y.Artist = y.Artist, x.Artist
			},
		})
	case "album":
		sort.Stable(custom{tracks, // 传入struct, 使用 sort.Stable 排序
			func(x, y *Track) bool {
				return x.Album < y.Album
			},
			func(x, y *Track) {
				x.Album, y.Album = y.Album, x.Album
			},
		})

	case "year":
		sort.Stable(custom{tracks, // 传入struct, 使用 sort.Stable 排序
			func(x, y *Track) bool {
				return x.Year < y.Year
			},
			func(x, y *Track) {
				x.Year, y.Year = y.Year, x.Year
			},
		})
	case "length":
		sort.Stable(custom{tracks, // 传入struct, 使用 sort.Stable 排序
			func(x, y *Track) bool {
				return x.Length < y.Length
			},
			func(x, y *Track) {
				x.Length, y.Length = y.Length, x.Length
			},
		})
	}
}

func main() {
	// 按照各个键独自排序
	fmt.Println("click before:")
	printTracks(tracks)
	fmt.Println("\nclick after:")
	
	colKeys := []string{"title", "artist", "album", "year", "length"}
	printInfo(colKeys)

}

func printInfo(keys []string) {
	for _, s := range keys {
		fmt.Printf("\nclick %s:\n", s)
		click(s)
		printTracks(tracks)
	}
}
