package webrepl

import (
	"context"
	"fmt"
	"waiig/lexer"
	"waiig/token"
	"waiig/web/view"

	"github.com/labstack/echo/v4"
)

func HandleEvaluate(c echo.Context) error {
	l := lexer.New(c.FormValue("in"))

	out := ""
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		out += fmt.Sprintf("%+v\n", tok)
	}
	return view.Print(out).Render(context.Background(), c.Response())
}

func HandleIndex(c echo.Context) error {
	return view.Index().Render(context.Background(), c.Response())
}
