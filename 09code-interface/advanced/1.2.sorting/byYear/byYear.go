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

// 实现 sort 接口, 按照 Year 排序
type byYear []*Track

func (b byYear) Len() int {
	return len(b)
}

// 比较前后两个元素
func (b byYear) Less(i, j int) bool {
	return b[i].Year < b[j].Year
}

func (b byYear) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	fmt.Println("sort Year:")
	sort.Sort(byYear(tracks))	// 转换tracks slice 成byYear 实现sort.Interface
	printTracks(tracks)

	fmt.Println("\nreverse sort Year:")
	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)
}
