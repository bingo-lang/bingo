package repl

import (
	"bufio"
	"fmt"
	"github.com/bingo-lang/bingo/evaluator"
	"github.com/bingo-lang/bingo/parser"
	"io"
	"strings"
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
		source := strings.NewReader(line)
		parser := parser.New(source)
		if program, err := parser.ParseProgram(); err == nil {
			for _, statement := range program.Statements {
				fmt.Println(evaluator.Eval(statement))
			}
		} else {
			fmt.Printf("Syntax Error: %s\n", err)
		}
	}
}
