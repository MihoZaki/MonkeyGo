package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MihoZaki/MonkeyGo/token"

	"github.com/MihoZaki/MonkeyGo/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		lines := scanner.Text()
		l := lexer.New(lines)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)

		}
	}

}
