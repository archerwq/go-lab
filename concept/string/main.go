// 字符串是不可变的字节序列，可以包含任意数据，但一般用来包含人类可读的文本.
// 文本字符串通常被解释为采用UTF8编码的Unicode码点序列(rune).
// 不可变意味着两个字符串可以共享相同的底层数据，这使得复制任何字符串、切片操作的代价是廉价的，都没有必要分配新的内存.
// Go语言源文件以及Go语言的文本字符串都是用UTF8编码的.
// Unicode字符集(码点，rune)，UTF8编码(一个码点编码成1到4个字节)，得益于UTF8编码的优良设计，诸多字符串操作都不需要解码操作.
// More about Golang String: https://blog.golang.org/strings
// To summarize, strings can contain arbitrary bytes, but when constructed from string literals,
// those bytes are (almost always, as long as they have no byte-level escapes) UTF-8.
// In Go, a character is represented by a value in single quote AKA character literal.
package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// 原生字符串面值，没有转义，全部内容都是字面值，包括换行
const usage = `Go is a tool for managing go source.

Usage:
	go command [arguments]
...
`

// 字符串面值
const hello = "Hello 世界"

func main() {
	showStringWithBytes()

	fmt.Print(usage)

	fmt.Println(len(hello))                    // 12, len返回字符串的字节数，不是rune字符数
	fmt.Println(utf8.RuneCountInString(hello)) // 8
	fmt.Printf("%c %[1]d\n", hello[7])         // 索引操作返回第i个字节的字节值，而不是第i个rune字符

	// 子字符串基于原始字符串的第i个字节开始到第j个字节生成一个新的字符串
	// +操作将两个字符串链接构造一个新的字符串
	fmt.Println(hello[:5] + " World") // "Hello World"

	// 字符串比较是逐个字节比较
	fmt.Println("2" > "10") // true

	// UTF8解码，当有非法UTF8编码输入时，会返回一个特别的Unicode字符'\uFFFD'(打印出来是黑钻中间是白色问号)
	for i := 0; i < len(hello); {
		r, size := utf8.DecodeRuneInString(hello[i:])
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		i += size
	}

	// 好在range操作会自动隐式解码UTF8字符串，注意索引的步长会大于1个字节，所以这里i可能不是连续值
	for i, r := range hello {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	rs := []rune(hello) // 类型转换时会进行解码
	fmt.Printf("%x\n", rs)
	fmt.Println(string(rs)) // 类型转换时会进行编码

	fmt.Println(intsToString([]int{1, 2, 3, 4}))
}

func showStringWithBytes() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
