package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule interface {
	Match(text string) (bool, string)
}

type MatchRule struct {
	match string
}

func (r *MatchRule) Match(text string) (bool, string) {
	ok := strings.HasPrefix(text, r.match)
	rem := text[len(r.match):]
	return ok, rem
}

type RefRule struct {
	rs *RuleSet
	ref int
}

func (r *RefRule) Match(text string) (bool, string) {
	sr, ok := r.rs.Rule(r.ref)
	if !ok {
		fmt.Fprintf(os.Stderr, "WARN: no rule with id %d\n", r.ref)
		return false, text
	}
	return sr.Match(text)
}

type OrRule struct {
	rules []Rule
}

func (r *OrRule) Match(text string) (bool, string) {
	for _, r := range r.rules {
		if ok, rem := r.Match(text); ok {
			return ok, rem
		}
	}
	return false, text
}

type ListRule struct {
	rules []Rule
}

func (r *ListRule) Match(text string) (bool, string) {
	for _, r := range r.rules {
		ok, rem := r.Match(text)
		if !ok {
			return false, rem
		}
		text = rem
	}
	return true, text
}

type RuleSet struct {
	rules map[int]Rule
}

func NewRuleSet() *RuleSet {
	return &RuleSet{
		rules: make(map[int]Rule),
	}
}

func (rs *RuleSet) parseSingleRule(rulestr string) (Rule, error) {
	if strings.HasPrefix(rulestr, "\"") &&
	   strings.HasSuffix(rulestr, "\"") {
		r := &MatchRule{
			match: rulestr[1:len(rulestr)-1],
		}
		return r, nil
	}
	id, err := strconv.Atoi(rulestr)
	if err != nil {
		return nil, err
	}
	r := &RefRule{
		rs: rs,
		ref: id,
	}
	return r, nil
}

func (rs *RuleSet) parseRule(rulestr string) (Rule, error) {
	orparts := strings.Split(rulestr, " | ")
	if len(orparts) > 1 {
		or := &OrRule{
			rules: make([]Rule, len(orparts)),
		}
		for i, p := range orparts {
			sr, err := rs.parseRule(p)
			if err != nil {
				return nil, nil
			}
			or.rules[i] = sr
		}
		return or, nil
	}
	listparts := strings.Split(rulestr, " ")
	if len(listparts) == 1 {
		r, err := rs.parseSingleRule(rulestr)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	lr := &ListRule{
		rules: make([]Rule, len(listparts)),
	}
	for i, p := range listparts {
		r, err := rs.parseSingleRule(p)
		if err != nil {
			return nil, err
		}
		lr.rules[i] = r
	}
	return lr, nil
}

func (rs *RuleSet) ParseRule(line string) error {
	idrule := strings.SplitN(line, ": ", 2)
	id, err := strconv.Atoi(idrule[0])
	if err != nil {
		return err
	}
	rule, err := rs.parseRule(idrule[1])
	if err != nil {
		return err
	}
	rs.rules[id] = rule
	return nil
}

func (rs *RuleSet) Rule(ruleId int) (Rule, bool) {
	r, ok := rs.rules[ruleId]
	return r, ok
}

func (rs *RuleSet) IsMatch(ruleId int, text string) bool {
	r, ok := rs.rules[ruleId]
	if !ok {
		return false
	}
	ok, rem := r.Match(text)
	return ok && rem == ""
}

func main() {
	rs := NewRuleSet()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		if err := rs.ParseRule(s.Text()); err != nil {
			panic(err)
		}
	}
	c := 0
	for s.Scan() {
		if rs.IsMatch(0, s.Text()) {
			c++
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", c)
}
