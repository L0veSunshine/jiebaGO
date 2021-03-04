package finalseg

import (
	"encoding/gob"
	"fmt"
	"os"
)

type (
	HmmStart map[byte]float64
	HmmTran  map[byte]map[rune]float64
	Path     map[byte][]byte
)

const (
	MinFloat  = -3.14e100
	HanRegExp = "([\u4e00-\u9fd5]+)"
	SkipReg   = "([a-zA-Z0-9]+(?:\\.\\d+)?%?)"
)

var HmmState = []byte{'B', 'M', 'E', 'S'}

func preState() Path {
	prevStatus := map[byte][]byte{}
	prevStatus['B'] = []byte{'E', 'S'}
	prevStatus['M'] = []byte{'M', 'B'}
	prevStatus['S'] = []byte{'S', 'E'}
	prevStatus['E'] = []byte{'B', 'M'}
	return prevStatus
}

func loadModel() (start HmmStart, trans, emit HmmTran) {
	var probStart = HmmStart{}
	probStart['B'] = -0.26268660809250016
	probStart['E'] = -3.14e+100
	probStart['M'] = -3.14e+100
	probStart['S'] = -1.4652633398537678
	var probTrans = loadEach("trans.prob")
	var probEmit = loadEach("emit.prob")

	return probStart, probTrans, probEmit
}

func loadEach(file string) (ht HmmTran) {
	f, e := os.Open(file)
	defer func() {
		_ = f.Close()
	}()
	if e != nil {
		fmt.Println(e)
	}
	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(&ht); err != nil {
		fmt.Println(err)
	}
	return ht
}
