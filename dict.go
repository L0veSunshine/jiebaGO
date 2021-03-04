package jiebaGo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	DirPath = "C:/Users/Xuan/Desktop/dict.txt"
)

type (
	Item map[string]Value

	Value struct {
		freq int
		part string
	}

	PrefixDict struct {
		Item
		total int
	}
)

func (i Item) Has(key string) bool {
	if _, ok := i[key]; ok {
		return true
	} else {
		return false
	}
}

func (i Item) Count() int {
	return len(i)
}

func (i Item) Freq(key string) int {
	return i[key].freq
}

func (pd *PrefixDict) addWord(word string, freq int) {
	pd.Item[word] = Value{
		freq: freq,
	}
	pd.total += freq
}

func (i Item) DefaultFreq(key string, value int) int {
	v := i.Freq(key)
	if v != 0 {
		return v
	}
	return value
}

func (i Item) addRecur(key string) {
	var length = len(key)
	var idx = 0
	for idx < length {
		_, size := utf8.DecodeRuneInString(key[idx:])
		idx += size
		substr := key[:idx]
		if !i.Has(substr) {
			sv := &Value{
				freq: 0,
			}
			i[substr] = *sv
		}
	}
}

func load() *PrefixDict {
	f, e := os.Open(DirPath)
	if e != nil {
		fmt.Println(e)
	}
	defer func() {
		_ = f.Close()
	}()
	buf := bufio.NewReader(f)
	var con = &PrefixDict{
		Item: make(Item, 500000),
	}
	total := 0
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF && err != nil {
			break
		}
		var SLine = string(line)
		res := strings.Split(SLine, " ")
		key := res[0]
		n, err := strconv.Atoi(res[1])
		if e != nil {
			fmt.Println(err)
		}
		total += n
		v := Value{
			freq: n,
		}
		con.Item[key] = v
		con.addRecur(key)
	}
	con.total = total
	return con
}
