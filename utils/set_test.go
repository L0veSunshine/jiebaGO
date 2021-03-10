package utils

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add("我")
	s.Add("你好")
	s.Add("大家")
	s.Add(2)
	s.Add(1)
	fmt.Println(s.ToArray())
	fmt.Println(s.Has("我"))
	fmt.Println(s.Remove("你好"))
	fmt.Println(s.ToArray())
}
