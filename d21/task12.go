package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	all := make([]string, 0)
	possible := make(map[string][]string)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		is := make([]string, 0)
		ls := make([]string, 0)
		islabel := false
		for _, p := range p {
			if p == "(contains" {
				islabel = true
				continue
			}
			if islabel {
				ls = append(ls,
					strings.TrimSuffix(
						strings.TrimSuffix(
							p,
							")"),
						","))
			} else {
				is = append(is, p)
			}
		}
		all = append(all, is...)
		for _, l := range ls {
			pi, ok := possible[l]
			if !ok {
				possible[l] = is
				continue
			}
			npi := make([]string, 0)
			for _, i1 := range pi {
				for _, i2 := range is {
					if i1 == i2 {
						npi = append(npi, i1)
					}
				}
			}
			possible[l] = npi
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	c := 0
	for _, i := range all {
		match := false
		for _, pis := range possible {
			for _, pi := range pis {
				if pi == i {
					match = true
					break
				}
			}
			if match {
				break
			}
		}
		if !match {
			c++
		}
	}
	fmt.Printf("%d\n", c)

	assignedIs := make(map[string]bool)
	for i := 0; i < len(possible); i++ {
		for l, is := range possible {
			if len(is) == 1 {
				assignedIs[is[0]] = true
			} else {
				nis := make([]string, 0)
				for _, ig := range is {
					if _, ok := assignedIs[ig]; !ok {
						nis = append(nis, ig)
					}
				}
				possible[l] = nis
			}
		}
	}

	ls := make([]string, 0)
	for l, _ := range possible {
		ls = append(ls, l)
	}
	sort.Strings(ls)

	for i, l := range ls {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%s", possible[l][0])
	}
	fmt.Printf("\n")
}

