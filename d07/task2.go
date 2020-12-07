package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dep struct {
	Num  int
	Name string
}

var rules map[string][]Dep

func parseRule(l string) {
	p := strings.Split(l, ",")
	dep := make([]Dep, 0)
	for _, v := range p[1:] {
		if v == "no other" {
			continue
		}
		dp := strings.SplitN(v, " ", 2)
		i, err := strconv.Atoi(dp[0])
		if err != nil {
			panic(err)
		}
		dep = append(dep, Dep{
			Num:  i,
			Name: dp[1],
		})
	}
	rules[p[0]] = dep
}

func evalRule(name string) int {
	dep, ok := rules[name]
	if !ok {
		return 0
	}
	i := 0
	for _, v := range dep {
		i += v.Num
		i += evalRule(v.Name) * v.Num
	}
	return i
}

func main() {
	rules = make(map[string][]Dep)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parseRule(s.Text())
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	fmt.Println(evalRule("shiny gold"))
}
