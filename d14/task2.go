package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const ADDR_LEN = 36

var mask []rune
var addrs = make(map[int64]int64)

func process(cmd, val string) {
	if cmd == "mask" {
		if len(val) != ADDR_LEN {
			panic("invalid input")
		}
		mask = []rune(val)
		return
	}

	id, err := strconv.ParseInt(cmd[4:len(cmd)-1], 10, 64)
	if err != nil {
		panic(err)
	}
	vali, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}

	idRunes := []rune(fmt.Sprintf(fmt.Sprintf("%%0%ds", ADDR_LEN), strconv.FormatInt(id, 2)))
	pb := 0
	for _, v := range mask {
		if v == 'X' {
			pb++
		}
	}
	pl := int(math.Pow(2, float64(pb)))
	for p := 0; p < pl; p++ {
		var idRunes2 [ADDR_LEN]rune
		xv := []rune(fmt.Sprintf(fmt.Sprintf("%%0%ds", pb), strconv.FormatInt(int64(p), 2)))
		xi := 0
		for i := 0; i < ADDR_LEN; i++ {
			switch mask[i] {
			case '0':
				idRunes2[i] = idRunes[i]
			case '1':
				idRunes2[i] = '1'
			case 'X':
				idRunes2[i] = xv[xi]
				xi++
			}
		}
		id, err := strconv.ParseInt(string(idRunes2[0:]), 2, 64)
		if err != nil {
			panic(fmt.Sprintf("parse id runes %v", err))
		}
		addrs[id] = vali
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		p := strings.Split(s.Text(), " = ")
		process(p[0], p[1])
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	var sum int64
	for _, v := range addrs {
		sum += v
	}
	fmt.Printf("%d\n", sum)
}
