package utils

import (
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func ByteToStr(b []byte) string {
	byteSliceHead := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: byteSliceHead.Data,
		Len:  byteSliceHead.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func StrToByte(s string) []byte {
	strSliceHead := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytes := reflect.SliceHeader{
		Data: strSliceHead.Data,
		Len:  strSliceHead.Len,
		Cap:  strSliceHead.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bytes))
}

func StrToRune(s string) []rune {
	data := StrToByte(s)
	lens := len(data)
	runes := make([]rune, 0, 3*lens/5)
	var n = 0
	for n < lens {
		r, size := utf8.DecodeRune(data[n:])
		runes = append(runes, r)
		n += size
	}
	return runes
}

func RuneToStr(rs []rune) string {
	lens := len(rs)
	bytes := make([]byte, 0, lens*3)
	buf := make([]byte, 3)
	for i := 0; i < len(rs); i++ {
		n := utf8.EncodeRune(buf, rs[i])
		bytes = append(bytes, buf[:n]...)
	}
	return ByteToStr(bytes)
}
