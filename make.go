package shogi

import (
	"math/rand"

	"github.com/ogiekako/shogi/board"
)

func MakeRandom(r *rand.Rand) *board.Ban {
	ban := &board.Ban{}
	for i := 0; i < 2; i++ {
		ban.Col = append(ban.Col, &board.Ban_Col{Koma: make([]board.Ban_Koma, 9, 9)})
		for j := 0; j < 9; j++ {
			ban.Col[i].Koma[j] = random(r)
		}
	}
	ban.Col[0].Koma[8] = board.Ban_E
	return ban
}

func random(r *rand.Rand) board.Ban_Koma {
	ks := []board.Ban_Koma{
		board.Ban_F,
		board.Ban_K,
		board.Ban_C,
		board.Ban_S,
		board.Ban_N,
	}
	return ks[r.Intn(len(ks))]
}
