package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	From int
	To   int
}

func (r Rule) Check(value int) bool {
	return value >= r.From && value <= r.To
}

type FieldIdList []int

func (l FieldIdList) Len() int {
	return len(l)
}

func (l FieldIdList) Less(i, j int) bool {
	return len(matchingFields[l[i]]) < len(matchingFields[l[j]])
}

func (l FieldIdList) Swap(i, j int) {
	tmp := l[i]
	l[i] = l[j]
	l[j] = tmp
}

var rules map[string][]Rule
var matchingFields [][]string
var fieldIds map[string]int
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

func checkAnyRule(rules []Rule, value int) bool {
	for _, r := range rules {
		if r.Check(value) {
			return true
		}
	}
	return false
}

func filterValid() {
	vt := make([][]int, 0)
	for _, t := range nTickets {
		allValid := true
		for _, v := range t {
			valid := false
			for _, rs := range rules {
				if checkAnyRule(rs, v) {
					valid = true
					break
				}
			}
			if !valid {
				allValid = false
				break
			}
		}
		if allValid {
			vt = append(vt, t)
		}
	}
	nTickets = vt
}

func prepareMatchingFields() {
	matchingFields = make([][]string, len(myTicket))
	for i := 0; i < len(myTicket); i++ {
		matchingFields[i] = make([]string, len(rules))
		j := 0
		for k, _ := range rules {
			matchingFields[i][j] = k
			j++
		}
	}
}

func filterFields(fieldList []string, value int) []string {
	newFieldList := make([]string, 0)
	for _, field := range fieldList {
		if checkAnyRule(rules[field], value) {
			newFieldList = append(newFieldList, field)
		}
	}
	return newFieldList
}

func filterMatchingFields() {
	for _, t := range nTickets {
		for i, v := range t {
			matchingFields[i] = filterFields(matchingFields[i], v)
		}
	}
}

func assignFieldIds() {
	fieldIds = make(map[string]int)
	fieldIdList := make(FieldIdList, len(myTicket))
	for i := range fieldIdList {
		fieldIdList[i] = i
	}
	sort.Sort(fieldIdList)
	for _, v := range fieldIdList {
		for _, field := range matchingFields[v] {
			if _, ok := fieldIds[field]; ok {
				continue
			}
			fieldIds[field] = v
			break
		}
	}
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
	filterValid()
	prepareMatchingFields()
	filterMatchingFields()
	assignFieldIds()
	v := 1
	for field, id := range fieldIds {
		if strings.HasPrefix(field, "departure") {
			v = v * myTicket[id]
		}
	}
	fmt.Printf("%d\n", v)
}
