package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	From int
	To   int
}

var rules map[string][]Rule
var myTicket []int
var nTickets [][]int

func parseTicket(text string) []int {
	p := strings.Split(text, ",")
	t := make([]int, len(p))
	for i, v := range p {
		vi, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		t[i] = vi
	}
	return t
}

func parseRules(text string) []Rule {
	p := strings.Split(text, " or ")
	r := make([]Rule, len(p))
	for i, v := range p {
		ft := strings.Split(v, "-")
		from, err := strconv.Atoi(ft[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(ft[1])
		if err != nil {
			panic(err)
		}
		r[i] = Rule{
			From: from,
			To:   to,
		}
	}
	return r
}

func parseLine(line string) {
	p := strings.SplitN(line, ": ", 2)
	switch p[0] {
	case "yt":
		myTicket = parseTicket(p[1])
	case "nt":
		t := parseTicket(p[1])
		nTickets = append(nTickets, t)
	default:
		rules[strings.TrimSuffix(p[0], ":")] = parseRules(p[1])
	}
}

func findInvalid() int {
	iv := 0
	for _, t := range nTickets {
		for _, v := range t {
			valid := false
			for _, rs := range rules {
				for _, r := range rs {
					if v >= r.From && v <= r.To {
						valid = true
					}
				}
				if valid {
					break
				}
			}
			if !valid {
				iv += v
			}
		}
	}
	return iv
}

func main() {
	rules = make(map[string][]Rule)
	nTickets = make([][]int, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parseLine(s.Text())
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", findInvalid())
}
