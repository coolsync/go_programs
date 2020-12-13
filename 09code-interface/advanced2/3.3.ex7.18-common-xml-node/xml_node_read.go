package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Node interface {
	String() string	
}

// 存放文本 node
type CharData string

func (c CharData) String() string {
	return string(c)
}

// 存放 元素 node
type Element struct {
	Type     xml.Name   // "Title", "div"
	Attr     []xml.Attr // id=p1
	Children []Node     // 子节点集合
}

// 通过 String() method 获取 node all content
func (e *Element) String() string {
	var childrens, attrs string

	// 提取 属性
	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}

	// 提取 子节点
	for _, child := range e.Children {
		childrens += child.String() // 递归遍历leaf节点, text node contains html 转义char
	}

	return fmt.Sprintf("<%s%s>%s</%s>", e.Type.Local, attrs, childrens, e.Type.Local)
}

func parse(dec *xml.Decoder) (Node, error) {
	var stack []*Element // 存放元素

	for {
		tok, err := dec.Token()
		if err != nil {
			// if err == io.EOF {
			// 	break
			// }
			return nil, fmt.Errorf("dec token err: %v", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, e)
			}
			stack = append(stack, e)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stack) == 1 { // html node
				return stack[0], nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, CharData(tok))
			}
		}
	}
}

func main() {
	node, err := parse(xml.NewDecoder(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", node)	// call node String method
}
