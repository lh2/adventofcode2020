package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cup struct {
	Value int
	Next  *Cup
}

func shuffleCups(max int, cups *Cup) *Cup {
	take := cups.Next
	cups.Next = take.Next.Next.Next
	dest := cups.Value
	var destCup *Cup
	for destCup == nil {
		dest = dest - 1
		if dest < 1 {
			dest = max - dest
		}
		for cup := cups.Next; ; cup = cup.Next {
			if cup.Value == cups.Value {
				break
			}
			if cup.Value == dest {
				destCup = cup
				break
			}
		}
	}

	dcn := destCup.Next
	destCup.Next = take
	take.Next.Next.Next = dcn
	return cups.Next
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var maxcup int
	cupss := strings.Split(strings.TrimSpace(string(data)), "")
	var first *Cup
	var cups *Cup
	for _, cup := range cupss {
		ci, err := strconv.Atoi(cup)
		if err != nil {
			panic(err)
		}
		if cups == nil {
			cups = &Cup{ci, nil}
			first = cups
		} else {
			c := &Cup{ci, nil}
			cups.Next = c
			cups = c
		}

		if ci > maxcup {
			maxcup = ci
		}
	}
	cups.Next = first
	cups = first
	for i := 0; i < 100; i++ {
		cups = shuffleCups(maxcup, cups)
	}
	res := ""
	found1 := false

	for cup := cups; ; cup = cup.Next {
		if found1 && cup.Value == 1 {
			break
		}
		if cup.Value == 1 {
			found1 = true
			continue
		}
		if found1 {
			res += strconv.Itoa(cup.Value)
		}
	}
	fmt.Println(res)
}
