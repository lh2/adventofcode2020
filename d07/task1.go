package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rules map[string][]string

func parseRule(l string) {
	p := strings.Split(l, ",")
	rules[p[0]] = p[1:]
}

func evalRule(name string) bool {
	if name == "shiny gold" {
		return true
	}
	dep, ok := rules[name]
	if !ok {
		return false
	}
	for _, v := range dep {
		if evalRule(v) {
			return true
		}
	}
	return false
}

func evalRules() {
	count := 0
	for k, _ := range rules {
		if k != "shiny gold" && evalRule(k) {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	rules = make(map[string][]string)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parseRule(s.Text())
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	evalRules()
}
