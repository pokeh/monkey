package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pokeh/monkey/lexer"
	"github.com/pokeh/monkey/token"
)

const PROMPT = "⊂((*'⊥'*))⊃ >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// %v -> the value in a default format when printing structs,
			// the plus flag (%+v) adds field names
			fmt.Printf("%+v\n", tok)
		}
	}
}
