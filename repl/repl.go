package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/zkpranav/imonkey/token"
	"github.com/zkpranav/imonkey/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%s", PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Printf("%+v\n", tk)
		}
 	}
}