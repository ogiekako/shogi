package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ogiekako/shogi"
	"github.com/ogiekako/shogi/board"
)

var baseLimit = flag.Int64("base_limit", math.MaxInt64, "base limit")

func main() {
	flag.Parse()
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	b := &board.Ban{}
	proto.UnmarshalText(string(buf), b)
	shogi.FindLoop(b, *baseLimit)
}
