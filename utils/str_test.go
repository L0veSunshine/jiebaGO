package utils

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

var text = []string{"此应用包"}
var rs = []rune{22240, 20026, 23427, 20351, 29992, 20102, 26576, 20123,
	21463, 38480, 21046, 30340, 21151, 33021, 12290}

func TestStrToRune(t *testing.T) {
	fmt.Println(string(rs))
	fmt.Println(RuneToStr(rs))

}

func BenchmarkRTS1(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = string([]rune{22240, 20026, 23427, 20351})
	}
	fmt.Println(s)
}

func BenchmarkRTS2(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = RuneToStr([]rune{22240, 20026, 23427, 20351})
	}
	fmt.Println(s)
}

func BenchmarkSTR1(b *testing.B) {
	var rs []rune
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range text {
			rs = []rune(s)
		}

	}
	fmt.Println(rs)
}

func BenchmarkSTR2(b *testing.B) {
	var rs []rune
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range text {
			rs = StrToRune(s)
		}
	}
	fmt.Println(rs)
}

func TestName(t *testing.T) {
	for _, s := range text {
		fmt.Println(StrToRune(s))
		fmt.Println(string(StrToRune(s)))
	}
}
func TestName1(t *testing.T) {
	var a []int
	var j int
	a = append(a, 1, 2, 3, 4, 5, 6, 7)
	s := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	for i := 0; i < s.Len; i++ {
		offset := unsafe.Sizeof(j) * uintptr(i)
		b := (*int)(unsafe.Pointer(s.Data + offset))
		fmt.Println(*b)
	}
}

func TestRuneToStr(t *testing.T) {
	type args struct {
		rs []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{rs: rs}, want: "因为它使用了某些受限制的功能。"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneToStr(tt.args.rs); got != tt.want {
				t.Errorf("RuneToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
