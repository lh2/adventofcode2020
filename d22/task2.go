package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcGameId(cards1, cards2 []int) string {
	s := ""
	for _, c := range cards1 {
		s += strconv.Itoa(c) + ","
	}
	s += "\n"
	for _, c := range cards2 {
		s += strconv.Itoa(c) + ","
	}
	return s
}

func cloneCards(cs []int, max int) []int {
	cs2 := make([]int, max)
	for i, c := range cs {
		if i >= max {
			break
		}
		cs2[i] = c
	}
	return cs2
}

func play(cards1, cards2 []int) (int, []int, []int) {
	games := make(map[string]bool)
	for {
		if len(cards1) == 0 {
			return 2, cards1, cards2
		}
		if len(cards2) == 0 {
			return 1, cards1, cards2
		}

		gid := calcGameId(cards1, cards2)
		if games[gid] {
			return 1, cards1, cards2
		}
		games[gid] = true

		c1 := cards1[0]
		c2 := cards2[0]
		cards1 = cards1[1:]
		cards2 = cards2[1:]

		var w int
		if len(cards1) >= c1 && len(cards2) >= c2 {
			w, _, _ = play(cloneCards(cards1, c1), cloneCards(cards2, c2))
		} else {
			if c1 > c2 {
				w = 1
			} else {
				w = 2
			}
		}

		switch w {
		case 1:
			cards1 = append(cards1, c1, c2)
		case 2:
			cards2 = append(cards2, c2, c1)
		}
	}
}

func calcScore(cards []int) int {
	ts := 0
	for i, c := range cards {
		ts += c * (len(cards) - i)
	}
	return ts
}

func parseCards(line string) []int {
	p := strings.Split(line, ",")
	cards := make([]int, len(p))
	for i, c := range p {
		n, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		cards[i] = n
	}
	return cards
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("EOF")
	}
	cards1 := parseCards(s.Text())
	if !s.Scan() {
		panic("EOF")
	}
	cards2 := parseCards(s.Text())
	if err := s.Err(); err != nil {
		panic(err)
	}
	pw, p1, p2 := play(cards1, cards2)
	var score int
	switch pw {
	case 1:
		score = calcScore(p1)
	case 2:
		score = calcScore(p2)
	}
	fmt.Printf("%d\n", score)
}
