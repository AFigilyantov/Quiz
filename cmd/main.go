package main

import (
	"flag"
)

func main() {

	filePath := flag.String("filePath", "problems.csv", "path to file *.csv with questions")
	shuffle := flag.Bool("shuffle", false, "should the programme shuffle questions")

	flag.Parse()

	game := Game{}
	game.File = *filePath
	game.Shuffle = *shuffle

	game.Start()

}
