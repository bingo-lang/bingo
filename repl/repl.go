package repl

import (
	"bufio"
	"fmt"
	"github.com/bingo-lang/bingo/evaluator"
	"github.com/bingo-lang/bingo/object"
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
		program := parser.Parse()
		gotten, _ := evaluator.Eval(program.Statements[0]).(*object.Integer)
		fmt.Println(gotten.Value)
	}
}
