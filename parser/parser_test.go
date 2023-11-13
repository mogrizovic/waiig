package parser

import (
	"testing"
	"waiig/ast"
	"waiig/lexer"
	"waiig/token"
)

func TestLetStatement(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
		`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct{ expectedIdentifier string }{
		{expectedIdentifier: "x"},
		{expectedIdentifier: "y"},
		{expectedIdentifier: "foobar"},
	}

	for i, tt := range tests {
		s := program.Statements[i]

		if s.TokenLiteral() != "let" {
			t.Errorf("letStmt.TokenLiteral not let. got=%q", s.TokenLiteral())
			return
		}
		letStmt, ok := s.(*ast.LetStatement)
		if !ok {
			t.Errorf("Statement is not a LetStatement. got=%T", s)
			return
		}
		if letStmt.Name.Token.Type != token.IDENT {
			t.Errorf("Identificator token in IDENT. got=%s", letStmt.Name.Token.Type)
			return
		}
		if letStmt.Name.Value != tt.expectedIdentifier {
			t.Errorf("Identificator not '%s'. got=%s", tt.expectedIdentifier, letStmt.Name.Value)
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return x;
	return add(5, x);
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Errorf("Expected 3 statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected *ast.ReturnStatement, got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got=%s", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	} else {
		t.Errorf("Parser has %d errors.", len(p.errors))
		for _, e := range p.errors {
			t.Errorf("parser error: %q", e)
		}
		t.FailNow()
	}
}
