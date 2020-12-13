package main

func main() {

/*

	// A Name represents an XML name (Local) annotated with a name space identifier (Space).
	// In tokens returned by Decoder.Token,
	// the Space identifier is given as a canonical URL,
	// not the short prefix used in the document being parsed.

	// 名称表示带有名称空间标识符（Space）的XML名称（本地）。
	// 在Decoder.Token返回的tokens中，
	// 空间标识符以规范的URL形式给出，
	// 不是要解析的文档中使用的短前缀。

	type Name struct {
		Space, Local string
	}

	// An Attr represents an attribute in an XML element (Name=Value).
	//Attr表示XML元素（名称=值）中的属性。

	type Attr struct {
		Name  Name
		Value string
	}

	// A Token is an interface holding one of the token types:
	// StartElement, EndElement, CharData, Comment, ProcInst, or Directive.
	// Token是持有Token类型之一的接口：
	// StartElement，EndElement，CharData，Comment，ProcInst或Directive。

	type Token interface{}

	// A StartElement represents an XML start element.
	// StartElement代表XML开始元素。

	type StartElement struct {
		Name Name
		Attr []Attr
	}

	// An EndElement represents an XML end element.
	// EndElement表示XML结束元素。

	type EndElement struct {
		Name Name
	}

	// A CharData represents XML character data (raw text),
	// in which XML escape sequences have been replaced by the characters they represent.

	// CharData表示XML字符数据（原始文本），
	// 其中的XML转义序列已由其表示的字符替换。

	type CharData []byte

	// A Comment represents an XML comment of the form <!--comment-->.
	// The bytes do not include the <!-- and --> comment markers.

	// 注释表示<!--comment-->形式的XML注释。
	// 字节不包含<!-- and -->注释标签。

	type Comment []byte


	type Decoder struct {...}


	// NewDecoder creates a new XML parser reading from r.
	// If r does not implement io.ByteReader,
	// NewDecoder will do its own buffering.

	// NewDecoder创建一个从r读取的新XML解析器。
	// 如果r没有实现io.ByteReader，
	// NewDecoder将执行自己的缓冲

	func NewDecoder(r io.Reader) *Decoder {
		d := &Decoder{
			ns:       make(map[string]string),
			nextByte: -1,
			line:     1,
			Strict:   true,
		}
		d.switchToReader(r)
		return d
	}


	func (d *Decoder) Token() (Token, error) {}


	// Token returns the next XML token in the input stream.
	// At the end of the input stream, Token returns nil, io.EOF.
	// Token返回输入流中的下一个XMLToken。
	// 在输入流的末尾，Token返回nil，io.EOF。
	//
	// Slices of bytes in the returned token data refer to the
	// parser's internal buffer and remain valid only until the next
	// call to Token. To acquire a copy of the bytes, call CopyToken
	// or the token's Copy method.
	// 返回的Token数据中的字节切片是指
	// 解析器的内部缓冲区，仅在下一个缓冲区有效
	// 调用Token。要获取字节的副本，请调用CopyToken
	// 或Token的Copy方法。
	//
	// Token expands self-closing elements such as <br/>
	// into separate start and end elements returned by successive calls.
	// Token扩展了诸如<br/>之类的自闭元素
	// 分为连续调用返回的单独的开始和结束元素。
	//
	// Token guarantees that the StartElement and EndElement
	// tokens it returns are properly nested and matched:
	// if Token encounters an unexpected end element
	// or EOF before all expected end elements,
	// it will return an error.
	// Token可确保StartElement和EndElement
	// 返回的Token已正确嵌套和匹配：
	// 如果Token遇到意外的结束元素
	// 或EOF在所有预期的结束元素之前，
	// 它将返回错误。
	//
	// Token implements XML name spaces as described by
	// https://www.w3.org/TR/REC-xml-names/.  Each of the
	// Name structures contained in the Token has the Space
	// set to the URL identifying its name space when known.
	// If Token encounters an unrecognized name space prefix,
	// it uses the prefix as the Space rather than report an error.

	// Token实现XML名称空间，如
	// https://www.w3.org/TR/REC-xml-names/。每个
	// Token中包含的名称结构具有空格
	// 设置为URL，以在已知时标识其名称空间。
	// 如果Token遇到无法识别的名称空间前缀，
	// 它使用前缀作为空格，而不是报告错误。

	
	
	
	// Token返回输入流中的下一个XMLToken。
	// 在输入流的末尾，Token返回nil，io.EOF。
	//
	// 返回的Token数据中的字节切片是指
	// 解析器的内部缓冲区，仅在下一个缓冲区有效
	// 调用Token。要获取字节的副本，请调用CopyToken
	// 或Token的Copy方法。
	//
	// Token扩展了诸如<br/>之类的自闭元素
	// 分为连续调用返回的单独的开始和结束元素。
	//
	// Token可确保StartElement和EndElement
	// 返回的Token已正确嵌套和匹配：
	// 如果Token遇到意外的结束元素
	// 或EOF在所有预期的结束元素之前，
	// 它将返回错误。
	//
	// Token实现XML名称空间，如
	// https://www.w3.org/TR/REC-xml-names/。每个
	// Token中包含的名称结构具有空格
	// 设置为URL，以在已知时标识其名称空间。
	// 如果Token遇到无法识别的名称空间前缀，
	// 它使用前缀作为空格，而不是报告错误。




	func (d *Decoder) Token() (Token, error) {
		var t Token
		var err error
		if d.stk != nil && d.stk.kind == stkEOF {
			return nil, io.EOF
		}
		if d.nextToken != nil {
			t = d.nextToken
			d.nextToken = nil
		} else if t, err = d.rawToken(); err != nil {
			switch {
			case err == io.EOF && d.t != nil:
				err = nil
			case err == io.EOF && d.stk != nil && d.stk.kind != stkEOF:
				err = d.syntaxError("unexpected EOF")
			}
			return t, err
	}

	if !d.Strict {
		if t1, ok := d.autoClose(t); ok {
			d.nextToken = t
			t = t1
		}
	}
	switch t1 := t.(type) {
	case StartElement:
		// In XML name spaces, the translations listed in the
		// attributes apply to the element name and
		// to the other attribute names, so process
		// the translations first.
		
		//在XML名称空间中,
		//属性适用于元素名称和
		//为其他属性名称，因此进行处理
		//首先翻译。

		for _, a := range t1.Attr {
			if a.Name.Space == xmlnsPrefix {
				v, ok := d.ns[a.Name.Local]
				d.pushNs(a.Name.Local, v, ok)
				d.ns[a.Name.Local] = a.Value
			}
			if a.Name.Space == "" && a.Name.Local == xmlnsPrefix {
				// Default space for untagged names
				v, ok := d.ns[""]
				d.pushNs("", v, ok)
				d.ns[""] = a.Value
			}
		}

		d.translate(&t1.Name, true)
		for i := range t1.Attr {
			d.translate(&t1.Attr[i].Name, false)
		}
		d.pushElement(t1.Name)
		t = t1

	case EndElement:
		d.translate(&t1.Name, true)
		if !d.popElement(&t1) {
			return nil, d.err
		}
		t = t1
	}
	return t, err
	}



	// A Decoder represents an XML parser reading a particular input stream.
	// The parser assumes that its input is encoded in UTF-8.
	// 解码器表示读取特定输入流的XML解析器。
	// 解析器假定其输入以UTF-8编码。

	type Decoder struct {
	// Strict defaults to true, enforcing the requirements
	// of the XML specification.
	// If set to false, the parser allows input containing common
	// mistakes:
	//	* If an element is missing an end tag, the parser invents
	//	  end tags as necessary to keep the return values from Token
	//	  properly balanced.
	//	* In attribute values and character data, unknown or malformed
	//	  character entities (sequences beginning with &) are left alone.
	//

	// 严格默认为true，以强制执行要求
	// XML规范。
	// 如果设置为false，则解析器允许包含common的输入
	// 错误：
	//  * 如果元素缺少结束标记，则解析器会发明
	// 结束标记以保持Token的返回值
	// 适当平衡。
	//  * 在属性值和字符数据中，未知或格式错误
	// 字符实体（以＆开头的序列）不予处理。


	// Setting:
	//
	//	d.Strict = false
	//	d.AutoClose = xml.HTMLAutoClose
	//	d.Entity = xml.HTMLEntity
	//
	// creates a parser that can handle typical HTML.
	//
	// Strict mode does not enforce the requirements of the XML name spaces TR.
	// In particular it does not reject name space tags using undefined prefixes.
	// Such tags are recorded with the unknown prefix as the name space URL.
	// 创建一个可以处理典型HTML的解析器。
	//
	// 严格模式不强制执行XML命名空间TR的要求。
	// 特别是它不会拒绝使用未定义前缀的名称空间标签。
	// 这样的标签以未知前缀记录为名称空间URL。


	Strict bool

	// When Strict == false, AutoClose indicates a set of elements to
	// consider closed immediately after they are opened, regardless
	// of whether an end element is present.
	//当Strict == false时，AutoClose指示要设置的一组元素
	//在打开后立即考虑关闭，无论
	//是否存在结束元素。

	AutoClose []string

	// Entity can be used to map non-standard entity names to string replacements.
	// The parser behaves as if these standard mappings are present in the map,
	// regardless of the actual map content:
	//实体可用于将非标准实体名称映射为字符串替换。
	//解析器的行为就像这些标准映射存在于映射中一样，
	//无论实际的地图内容如何：

	//
	//	"lt": "<",
	//	"gt": ">",
	//	"amp": "&",
	//	"apos": "'",
	//	"quot": `"`,

	Entity map[string]string

	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// non-UTF-8 charset into UTF-8. If CharsetReader is nil or
	// returns an error, parsing stops with an error. One of the
	// CharsetReader's result values must be non-nil.

	// CharsetReader，如果不为nil，则定义要生成的function
	// charset-conversion读取器，从提供的内容进行转换
	//将非UTF-8字符集转换为UTF-8。如果CharsetReader为nil或
	//返回错误，由于错误而停止解析。其中一个
	// CharsetReader的结果值必须为非零。

	CharsetReader func(charset string, input io.Reader) (io.Reader, error)

	// DefaultSpace sets the default name space used for unadorned tags,
	// as if the entire XML stream were wrapped in an element containing
	// the attribute xmlns="DefaultSpace".

	// DefaultSpace设置用于无修饰标记的默认名称空间，
	//好像整个XML流都包装在包含
	//属性xmlns =“ DefaultSpace”。

	DefaultSpace string

	r              io.ByteReader
	t              TokenReader
	buf            bytes.Buffer
	saved          *bytes.Buffer
	stk            *stack
	free           *stack
	needClose      bool
	toClose        Name
	nextToken      Token
	nextByte       int
	ns             map[string]string
	err            error
	line           int
	offset         int64
	unmarshalDepth int
	}
*/
}
