package main

import (
	"fmt"
	"github.com/bingo-lang/bingo/repl"
	"os"
)

func main() {
	fmt.Printf("Bingo 0.0.1 [WIP]\n\n")
	fmt.Println("Press Ctrl + C to exit.")
	repl.Start(os.Stdin, os.Stdout)
}
