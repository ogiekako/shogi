package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ogiekako/shogi"
)

func main() {
	r := newRand()
	b := shogi.MakeRandom(r)
	fmt.Println(proto.MarshalTextString(b))
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
