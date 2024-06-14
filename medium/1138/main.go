package main

import (
	"log"
	"strings"
)

type pos struct {
	x int
	y int
}

func generateBoard() map[byte]pos {
	board := []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}
	m := make(map[byte]pos, 26)
	for i := 0; i < 6; i++ {
		for j := 0; j < len(board[i]); j++ {
			m[board[i][j]] = pos{
				x: j,
				y: i,
			}
		}
	}
	return m
}

func outOfBoundaries(p pos) bool {
	if p.y == 5 {
		return p.x != 0
	}
	return p.x > 5 || p.y > 5
}

const (
	//ignoreUp
	ignoreDown = 1 << iota
	ignoreRight
)

func generatePath(from pos, to pos, flag int) string {
	res := ""
	for from.x != to.x || from.y != to.y {
		switch {
		case from.y > to.y:
			res += strings.Repeat("U", from.y-to.y)
			from.y = to.y
		case from.y < to.y && (flag&ignoreDown) == 0:
			if outOfBoundaries(pos{x: from.x, y: to.y}) {
				if 4-to.y <= 0 {
					return res + generatePath(from, to, flag|ignoreDown)
				}
			}
			res += strings.Repeat("D", to.y-from.y)
			from.y = to.y
		case from.x < to.x && (flag&ignoreRight) == 0:
			if outOfBoundaries(pos{x: to.x, y: from.y}) {
				if 4-to.x <= 0 {
					return res + generatePath(from, to, flag|ignoreDown)
				}
			}
			res += strings.Repeat("R", to.x-from.x)
			from.x = to.x
		case from.x > to.x:
			res += strings.Repeat("L", from.x-to.x)
			from.x = to.x
		}
		flag = 0
	}
	return res + "!"
}

func alphabetBoardPath(target string) string {
	board := generateBoard()
	p := pos{}
	res := ""
	for i := range target {
		res += generatePath(p, board[target[i]], 0)
		p = board[target[i]]
	}
	return res
}

type test struct {
	target   string
	expected string
}

func main() {
	tests := []test{
		{target: "leet", expected: "DDR!UURRR!!DDD!"},
		{target: "code", expected: "RR!DDRR!UUL!R!"},
		{target: "zb", expected: "DDDDD!UUUUUR!"},
		{target: "zdz", expected: "DDDDD!UUUUURRR!DDDDLLLD!"},
		{target: "yz", expected: "DDDDRRRR!LLLLD!"},
	}
	for _, t := range tests {
		res := alphabetBoardPath(t.target)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %v, got %v", t.target, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.target)
		}
	}
}
