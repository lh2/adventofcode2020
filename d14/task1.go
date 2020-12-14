package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ADDR_LEN = 36

var mask []rune
var addrs = make(map[int][ADDR_LEN]rune)

func process(cmd, val string) {
	if cmd == "mask" {
		if len(val) != ADDR_LEN {
			panic("invalid input")
		}
		mask = []rune(val)
		return
	}

	id, err := strconv.Atoi(cmd[4 : len(cmd)-1])
	if err != nil {
		panic(err)
	}
	vali, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}

	runes := []rune(fmt.Sprintf(fmt.Sprintf("%%0%ds", ADDR_LEN), strconv.FormatInt(vali, 2)))
	var addr [ADDR_LEN]rune
	if a, ok := addrs[id]; ok {
		addr = a
	}

	for i := 0; i < ADDR_LEN; i++ {
		if mask[i] != 'X' {
			addr[i] = mask[i]
		} else {
			addr[i] = runes[i]
		}
	}

	addrs[id] = addr
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
		iv, err := strconv.ParseInt(string(v[0:]), 2, 64)
		if err != nil {
			panic(err)
		}
		sum += iv
	}
	fmt.Printf("%d\n", sum)
}
