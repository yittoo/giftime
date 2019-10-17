package main

import (
	"flag"
	"fmt"
	"image/gif"
	"log"
	"os"
)

var fn *string

func init() {
	fn = flag.String("i", "", "gif name to calculate gif length of")
}

func main() {
	f, err := os.Open(*fn)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	g, err := gif.DecodeAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	var l int
	for _, v := range g.Delay {
		l += v
	}
	fmt.Println(l * 10)
}
