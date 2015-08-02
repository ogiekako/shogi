package main

import (
	"flag"
	"github.com/ogiekako/shogi"
	"log"
	"math/rand"
	"os"
	_ "time"
)

var (
	baseLimit = flag.Int64("base_limit", 63356, "base limit")
	objStep   = flag.Int64("obj_step", int64(87635520), "objective step")
)

func main() {
	r := rand.New(rand.NewSource(99))
	max := int64(10000)
	for {
		ban := shogi.MakeRandom(r)
		step, problem, f := shogi.FindLoop(ban, *baseLimit)
		if step > max {
			max = step
			log.Printf("New record: %d  problem: %s  file: %s\n", max, shogi.Sshow(problem), f.Name())

			if max > *objStep {
				break
			}
		} else {
			err := os.Remove(f.Name())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
