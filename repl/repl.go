package repl

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(stdin io.Reader, stdout io.Writer) {
	input := bufio.NewScanner(stdin)
	for {
		fmt.Printf(PROMPT)
		scanned := input.Scan()
		if !scanned {
			//something went wrong
			return
		}
		line := input.Text()
		fmt.Println(line)
	}
}
