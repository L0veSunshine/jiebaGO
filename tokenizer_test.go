package jiebaGo

import (
	"fmt"
	"testing"
	"xuan/jiebaGo/utils"
)

func TestTokenizer_Tokenize(t *testing.T) {
	token := NewTokenizer()
	a := token.Cut("此应用包不支持通过应用安装程序安装，因为它使用了某些受限制的功能。", false)
	b := token.Tokenize("此应用包不支持通过应用安装程序安装，因为它使用了某些受限制的功能。", true, true)
	fmt.Println(a, b.ToJson())
}

func BenchmarkBToS1(b *testing.B) {
	t := []byte("此应用包不支持通过应用安装程序安装，因为它使用了某些受限制的功能。")
	var a string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = utils.ByteToStr(t)
	}
	fmt.Println(a)
}

func BenchmarkBToS2(b *testing.B) {
	t := []byte("此应用包不支持通过应用安装程序安装，因为它使用了某些受限制的功能。")
	var a string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = string(t)
	}
	fmt.Println(a)
}

func BenchmarkSToB1(b *testing.B) {
	t := "应用"
	var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = utils.StrToByte(t)
	}
	fmt.Println(bs)
}

func BenchmarkSToB2(b *testing.B) {
	t := "应用"
	var bs []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = []byte(t)
	}
	fmt.Println(bs)
	fmt.Println([]rune(t))
}
