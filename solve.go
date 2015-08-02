package shogi

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ogiekako/shogi/board"
)

const bucket = 1 << 20

func FindLoop(ban *board.Ban, baseLimit int64) (int64, *board.Ban, *os.File) {
	for limit := baseLimit; ; limit *= 2 {
		f, err := ioutil.TempFile("/tmp", "eigou")
		if err != nil {
			log.Fatalln(err)
		}
		var step int64
		step, ban = solve(ban, limit, f)
		if step >= 0 {
			log.Printf("output: %s\n", f.Name())
			log.Printf("step: %d\n", step*2)
			return step, ban, f
		}
		err = os.Remove(f.Name())
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (b *Ban) get(x, y int) board.Ban_Koma {
	x, y = adjust(x, y)
	return b.ban.Col[x].Koma[y]
}

type Ban struct {
	ban *board.Ban
	// Empty (x,y)
	x int
	y int
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

func (b *Ban) swap(x, y int) {
	k := b.ban.Col[b.x].Koma[b.y]
	b.ban.Col[b.x].Koma[b.y] = b.ban.Col[x].Koma[y]
	b.ban.Col[x].Koma[y] = k
	b.x, b.y = x, y
}

func empty(b *board.Ban) (x, y int) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 9; j++ {
			if b.Col[i].Koma[j] == board.Ban_E {
				return i, j
			}
		}
	}
	panic("No empty")
}

func (b *Ban) move() {
	x, y := b.x, b.y

	px, py := 0, 0
	switch b.get(x, y-1) {
	case board.Ban_F, board.Ban_K:
		px, py = adjust(x, y+1)
		// (1,5)
	case board.Ban_S:
		px, py = adjust(x+1, y+5)
		// (1,6)
	case board.Ban_C:
		px, py = adjust(x+1, y+6)
	case board.Ban_N:
		px, py = adjust(x+1, y+2)
	}
	b.swap(px, py)
}

func newBan(b *board.Ban) *Ban {
	x, y := empty(b)
	return &Ban{
		ban: b,
		x:   x,
		y:   y,
	}
}

// solve returns step or the board moved limit times.
func solve(init *board.Ban, limit int64, output *os.File) (int64, *board.Ban) {
	w := gzip.NewWriter(output)
	defer w.Close()
	log.Printf("solving: %s limit: %d", Sshow(init), limit)
	ban := newBan(proto.Clone(init).(*board.Ban))
	for step := int64(0); step < limit; {
		fmt.Fprintf(w, "%d%d\n", ban.x+1, ban.y+1)
		ban.move()
		step++

		if proto.Equal(init, ban.ban) {
			return step, ban.ban
		}
	}
	return -1, ban.ban
}

func Sshow(b *board.Ban) string {
	res := ""
	for j := 0; j < 9; j++ {
		for i := 1; i >= 0; i-- {
			res += b.Col[i].Koma[j].String()
		}
		res += " "
	}
	return res
}
