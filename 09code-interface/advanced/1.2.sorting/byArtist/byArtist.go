package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/*
func (*Writer) Init
func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
初始化一个Writer，第一个参数指定格式化后的输出目标，其余的参数控制格式化：

minwidth 最小单元长度
tabwidth tab字符的宽度
padding  计算单元宽度时会额外加上它
padchar  用于填充的ASCII字符，
         如果是'\t'，则Writer会假设tabwidth作为输出中tab的宽度，且单元必然左对齐。
flags    格式化控制
*/

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

// 将字符串时间格式 转换为 time 格式
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// 按照table格式打印输出格式
func printTracks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\t\n"
	tbw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 0, '\t', 0)         // 初始化 Writer输出格式
	fmt.Fprintf(tbw, format, "Title", "Artist", "Album", "Year", "Length") // tabw实现接口 io.Writer
	fmt.Fprintf(tbw, format, "------", "------", "------", "------", "------")

	for i := range tracks {
		fmt.Fprintf(tbw, format, tracks[i].Title, tracks[i].Artist, tracks[i].Album, tracks[i].Year, tracks[i].Length)
	}
	tbw.Flush()
}

// 实现sort接口, 按照 Artist 排序
type byArtist []*Track

func (b byArtist) Len() int {
	return len(b)
}

func (b byArtist) Less(i, j int) bool {
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	fmt.Println("byArtist:")

	sort.Sort(byArtist(tracks)) // 将 tracks 转换成 byArtist类型, 实现 sort.Interface
	printTracks(tracks)

	// 反转 排序
	fmt.Println("\nreverse byArtist:")

	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
}
