package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ogiekako/shogi/board"
)

func main() {
	ban := &board.Ban{}
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	proto.UnmarshalText(string(buf), ban)

	step := solve(ban)
	fmt.Println(step)
}

var memo = make(map[string]bool)

// add returns true if memo did not already contain the specified element
func add(b *board.Ban) bool {
	buf, err := proto.Marshal(b)
	if err != nil {
		log.Fatalln(err)
	}
	if key := string(buf); memo[key] {
		return false
	} else {
		memo[key] = true
		return true
	}
}

func get(b *board.Ban, x, y int) board.Ban_Koma {
	return b.Col[x].Koma[y]
}

func adjust(x, y int) (int, int) {
	if x < 0 {
		x += 2
	} else if x >= 2 {
		x -= 2
	}
	if y < 0 {
		y += 9
	} else if y >= 9 {
		y -= 9
	}
	return x, y
}

func swaped(b *board.Ban, x1, y1, x2, y2 int) *board.Ban {
	b = proto.Clone(b).(*board.Ban)
	k := b.Col[x1].Koma[y1]
	b.Col[x1].Koma[y1] = b.Col[x2].Koma[y2]
	b.Col[x2].Koma[y2] = k
	return b
}

func nexts(b *board.Ban) []*board.Ban {
	x, y := 0, 0
loop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 9; j++ {
			if b.Col[i].Koma[j] == board.Ban_E {
				x, y = i, j
				break loop
			}
		}
	}
	var res []*board.Ban
	ox, oy := adjust(x, y-2)
	if k := get(b, ox, oy); k == board.Ban_K || k == board.Ban_F {
		nx, ny := adjust(ox, oy+1)
		res = append(res, swaped(b, x, y, nx, ny))
	}
	ox, oy = adjust(x+1, y-6)
	if k := get(b, ox, oy); k == board.Ban_S {
		nx, ny := adjust(ox, oy+1)
		res = append(res, swaped(b, x, y, nx, ny))
	}
	ox, oy = adjust(x+1, y-7)
	if k := get(b, ox, oy); k == board.Ban_C {
		nx, ny := adjust(ox, oy+1)
		res = append(res, swaped(b, x, y, nx, ny))
	}
	return res
}

func solve(init *board.Ban) int {
	fmt.Printf("solving:\n%v\n", sshow(init))
	// add(init)
	var bq []*board.Ban
	var sq []int
	bq = append(bq, init)
	sq = append(sq, 0)
	for max := 0; ; {
		b := bq[0]
		s := sq[0]

		if max < s {
			max = s
			if max%10000 == 0 {
				if max%100000 == 0 {
					fmt.Printf("\n%d", max)
				}
				fmt.Printf(".")
			}
		}

		bq = bq[1:]
		sq = sq[1:]
		for _, nb := range nexts(b) {
			if proto.Equal(init, nb) {
				fmt.Println()
				return s + 1
			}
			// if !add(nb) {
			// continue
			// }
			bq = append(bq, nb)
			sq = append(sq, s+1)
		}
	}
}

func sshow(b *board.Ban) string {
	res := ""
	for j := 0; j < 9; j++ {
		for i := 1; i >= 0; i-- {
			res += b.Col[i].Koma[j].String()
		}
		res += "\n"
	}
	return res
}
