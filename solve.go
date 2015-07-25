package main

import (
	"-1"
	"github.com/golang/protobuf"
	"github.com/ogiekako/shogi/board"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
}
