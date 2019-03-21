package main

import "fmt"
import "strings"

func isSlash(r rune) bool {
	return r == '\\' || r == '/'
}


func Slash(r rune) rune {
	if r == '\\' {
		return '/'
	}
	return r
}

func main() {
	s := "hello world,你好世界!!!"
	fmt.Println(s)
	//统计!重复次数
	fmt.Println(strings.Count(s, "!"))
	//判断abc是否在字符串中
	fmt.Println(strings.Contains(s, "abc"))
	//判断字符串中是否包含任意一个字符
	fmt.Println(strings.ContainsAny(s, "abe"))
	//判断字符串中是否包含rune类型的字符,用单引号
	fmt.Println(strings.ContainsRune(s, '好'))
	//返回指定子串索引，没有返回-1
	fmt.Println(strings.Index(s, "o"))
	//返回指定子串最后出现的索引，没有返回-1
	fmt.Println(strings.LastIndex(s, "o"))
	//返回rune类型第一次出现的索引
	fmt.Println(strings.IndexRune(s, '好'))
	//返回指定串中任意字符在s中第一次出现的索引
	fmt.Println(strings.IndexAny(s, "hed"))
	//返回指定串中任意字符在s中最后一次出现的索引
	fmt.Println(strings.LastIndexAny(s, "hed"))
	//split转为数组
	//func SplitN(s, sep string, n int) []string
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表。
	// 如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
	// 参数 n 表示最多切分出几个子串，超出的部分将不再切分。
	// 如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分
	fmt.Printf("%q\n", strings.SplitN(s, ",",2))
	// SplitAfterN 以 sep 为分隔符，将 s 切分成多个子串，结果中包含 sep 本身
	fmt.Printf("%q\n", strings.SplitAfterN(s, ",",2))
	// Split 以 sep 为分隔符，将 s 切分成多个子切片，结果中不包含 sep 本身
	//func Split(s, sep string) []string
	fmt.Printf("%q\n", strings.Split(s, ","))
	// SplitAfter 以 sep 为分隔符，将 s 切分成多个子切片，结果中包含 sep 本身
	fmt.Printf("%q\n", strings.SplitAfter(s, ","))
	// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
	// 空白字符有：\t, \n, \v, \f, \r, ' ', U+0085 (NEL), U+00A0 (NBSP)
	// 如果 s 中只包含空白字符，则返回一个空列表
	//func Fields(s string) []string
	fmt.Printf("%q\n", strings.Fields(s))
	// FieldsFunc 以一个或多个满足 f(rune) 的字符为分隔符，
	// 将 s 切分成多个子串，结果中不包含分隔符本身。
	// 如果 s 中没有满足 f(rune) 的字符，则返回一个空列表。
	//func FieldsFunc(s string, f func(rune) bool) []string
	s = "C:\\Windows\\System32\\FileName"
	fmt.Printf("%q\n", strings.FieldsFunc(s, isSlash))
	//数组切片转字符串
	ss := []string{"Monday", "Tuesday", "Wednesday"}
	fmt.Println(strings.Join(ss, ","))
	//HasPrefix 判断字符串 s 是否以 prefix 开头
	//func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix(s, "hello"))
	// HasSuffix 判断字符串 s 是否以 prefix 结尾
	//func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix(s, "hello"))
	// Map 将 s 中满足 mapping(rune) 的字符替换为 mapping(rune) 的返回值。
	// 如果 mapping(rune) 返回负数，则相应的字符将被删除。
	//func Map(mapping func(rune) rune, s string) string
	s = "C:\\Windows\\System32\\FileName"
	fmt.Println(strings.Map(Slash, s))
	// Repeat 将 count 个字符串 s 连接成一个新的字符串
	//func Repeat(s string, count int) string
	fmt.Printf("%q\n", strings.Repeat(s, 3))
	// ToUpper 将 s 中的所有字符修改为其大写格式
	//func ToUpper(s string) string
	s = "hello"
	fmt.Println(strings.ToUpper(s))
	// ToLower 将 s 中的所有字符修改为其小写格式
	s = "HELLO"
	fmt.Println(strings.ToLower(s))
	// Title 将 s 中的所有单词的首字母修改为其 Title 格式
	s = "hello world"
	fmt.Println(strings.Title(s))
	// Trim 将删除 s 首尾连续的包含在 cutset 中的字符
	//func Trim(s string, cutset string) string
	s = "  hello word  "
	fmt.Printf("%q\n", strings.Trim(s, " "))
	// TrimLeft 将删除 s 头部连续的包含在 cutset 中的字符
	s = "  hello word  "
	fmt.Printf("%q\n", strings.TrimLeft(s, " "))
	// TrimRight 将删除 s 尾部连续的包含在 cutset 中的字符
	s = "  hello word  "
	fmt.Printf("%q\n", strings.TrimRight(s, " "))
	// TrimSpace 将删除 s 首尾连续的的空白字符
	s = "  hello word  "
	fmt.Printf("%q\n", strings.TrimSpace(s))
	// TrimPrefix 删除 s 头部的 prefix 字符串
	// 如果 s 不是以 prefix 开头，则返回原始 s
	//func TrimPrefix(s, prefix string) string
	s = "!!!hello world !!!"
	fmt.Printf("%q\n", strings.TrimPrefix(s, "!!!"))
	// TrimSuffix 删除 s 尾部的 suffix 字符串
	//TrimSuffix只是去掉s字符串结尾的suffix字符串，只是去掉１次，而TrimRight是一直去掉s字符串右边的字符
	s = "!!!hello world !!!"
	fmt.Printf("%q\n", strings.TrimSuffix(s, "!!!"))
	// Replace 返回 s 的副本，并将副本中的 old 字符串替换为 new 字符串
	// 替换次数为 n 次，如果 n 为 -1，则全部替换
	// 如果 old 为空，则在副本的每个字符之间都插入一个 new
	//func Replace(s, old, new string, n int) string
	s = "hello 世界"
	fmt.Println(strings.Replace(s, "世界", "world", -1))
	// EqualFold 判断 s 和 t 是否相等。忽略大小写，同时它还会对特殊字符进行转换
	// 比如将“ϕ”转换为“Φ”、将“Ǆ”转换为“ǅ”等，然后再进行比较
	//func EqualFold(s, t string) bool
	fmt.Println(strings.EqualFold("helloworld","HELLOWORLD"))


}
