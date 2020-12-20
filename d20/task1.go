package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TILE_SIZE = 10

type Point struct {
	X int
	Y int
}

type Corner int

const (
	_ Corner = 1 << iota
	Top
	Right
	Bottom
	Left
)

type Orientation struct {
	Corner  Corner
	Flipped bool
}

type Tile struct {
	ID  int
	Map map[Point]bool

	orientation Orientation

	Corners map[Orientation]int
}

func (t *Tile) HasCorner(id int) bool {
	for _, c := range t.Corners {
		if c == id {
			return true
		}
	}
	return false
}

func (t *Tile) SetOrientation(o Orientation) {
	t.orientation = o
}

func reserveString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (t *Tile) GenerateCornerCache() error {
	t.Corners = make(map[Orientation]int)
	for _, i := range []Corner{Top, Right, Bottom, Left} {
		x := 0
		y := 0
		if i == Right {
			x = TILE_SIZE - 1
		} else if i == Bottom {
			y = TILE_SIZE - 1
		}
		bitstr := ""
		for x < TILE_SIZE && y < TILE_SIZE {
			if b, ok := t.Map[Point{x, y}]; !ok {
				panic(fmt.Errorf("coordinate %d %d does not exist!\n", x, y))
			} else if b {
				bitstr += "1"
			} else {
				bitstr += "0"
			}

			if i == Top || i == Bottom {
				x++
			} else {
				y++
			}
		}

		flipped := false
		for j := 0; j < 2; j++ {
			cid, err := strconv.ParseInt(bitstr, 2, 32)
			if err != nil {
				return err
			}
			t.Corners[Orientation{
				Corner:  i,
				Flipped: flipped,
			}] = int(cid)
			bitstr = reserveString(bitstr)
			flipped = true
		}
	}
	return nil
}

func NewTile(id int) Tile {
	return Tile{
		ID:  id,
		Map: make(map[Point]bool),
	}
}

func main() {
	tiles := make([]Tile, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		idstr := strings.Split(s.Text(), " ")[1]
		idstr = idstr[:len(idstr)-1]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			panic(err)
		}
		tile := NewTile(id)
		y := 0
		for s.Scan() {
			line := s.Text()
			if line == "" {
				break
			}
			for x, v := range []rune(line) {
				tile.Map[Point{x, y}] = v == '#'
			}
			y++
		}
		if err := tile.GenerateCornerCache(); err != nil {
			panic(err)
		}
		tiles = append(tiles, tile)
	}

	cornerIDs := 1
	for _, t1 := range tiles {
		cc := 0
		for _, c := range t1.Corners {
			for _, t2 := range tiles {
				if t1.ID == t2.ID {
					continue
				}
				if t2.HasCorner(c) {
					cc++
				}
			}
		}
		if cc == 4 {
			cornerIDs *= t1.ID
		}
	}
	fmt.Printf("%d\n", cornerIDs)
}
