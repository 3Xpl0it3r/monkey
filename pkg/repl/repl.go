package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/3Xpl0it3r/monkey/pkg/lexer"
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for true {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
