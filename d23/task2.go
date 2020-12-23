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

func shuffleCups(max int, cups *Cup, cupmap map[int]*Cup) *Cup {
	take := cups.Next
	cups.Next = take.Next.Next.Next

	dest := cups.Value
	for {
		dest = dest - 1
		if dest < 1 {
			dest = max - dest
		}
		if dest != take.Value &&
			dest != take.Next.Value &&
			dest != take.Next.Next.Value {
			break
		}
	}
	destCup := cupmap[dest]

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
	cupmap := make(map[int]*Cup)
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
		cupmap[ci] = cups

		if ci > maxcup {
			maxcup = ci
		}
	}
	for v := maxcup + 1; v <= 1_000_000; v++ {
		c := &Cup{v, nil}
		cups.Next = c
		cups = c
		cupmap[v] = cups
	}
	cups.Next = first
	cups = first
	for i := 0; i < 10_000_000; i++ {
		cups = shuffleCups(1_000_000, cups, cupmap)
	}
	res := cupmap[1].Next.Value * cupmap[1].Next.Next.Value
	fmt.Printf("%d\n", res)
}
