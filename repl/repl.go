package repl

import (
	"bufio"
	"fmt"
	"io"
	"waiig/lexer"
	"waiig/token"
)

const (
    PROMPT = ">> "
)

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)

        scanned := scanner.Scan()
        if !scanned {
            return
        }
        l := lexer.New(scanner.Text())

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
