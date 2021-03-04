package utils

import (
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func SliceStr(str string, start, end int) []byte {
	if start < 0 {
		start = 0
	}
	bStr := []byte(str)
	byteLen := len(bStr)
	runeLen := utf8.RuneCount(bStr)
	var i, j = 0, 0
	if end > runeLen || end < 0 {
		end = runeLen
		j = byteLen
	}
	if start > end {
		return nil
	}
	dis := start + runeLen - end
	if dis > runeLen/3 {
		var fIdx, bIdx = 0, byteLen
		for count := 1; fIdx < byteLen; count++ {
			_, size := utf8.DecodeRune(bStr[fIdx:])
			fIdx += size
			if count == start {
				if start != 0 {
					i = fIdx
					break
				}
			}
		}
		for count := runeLen - 1; bIdx > 0; count-- {
			_, size := utf8.DecodeLastRune(bStr[:bIdx])
			bIdx -= size
			if count == end {
				j = bIdx
				break
			}
		}
	} else {
		frontDis := start + end
		backDis := 2*runeLen - frontDis
		if frontDis <= backDis {
			var idx = 0
			for count := 1; idx < byteLen; count++ {
				_, size := utf8.DecodeRune(bStr[idx:])
				idx += size
				if count == start {
					if start != 0 {
						i = idx
					}
				}
				if count == end {
					j = idx
					break
				}
			}
		} else {
			var idx = byteLen
			for count := runeLen - 1; idx > 0; count-- {
				_, size := utf8.DecodeLastRune(bStr[:idx])
				idx -= size
				if count == start {
					if start != 0 {
						i = idx
						break
					}
				}
				if count == end {
					j = idx
				}
			}
		}
	}
	return bStr[i:j]
}

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
