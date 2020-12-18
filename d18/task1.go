package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

var UnexpectedEndOfInput = errors.New("Unexpected end of input")

type Expr interface {
	Eval() int
	String() string
}

type Group struct {
	inner Expr
}

func (g *Group) Eval() int {
	return g.inner.Eval()
}

type UnaryExpr struct {
	a int
}

func (e *UnaryExpr) Eval() int {
	return e.a
}

type BinaryExpr struct {
	a  Expr
	b  Expr
	op rune
}

func (e *BinaryExpr) Eval() int {
	switch e.op {
	case '+':
		return e.a.Eval() + e.b.Eval()
	case '*':
		return e.a.Eval() * e.b.Eval()
	}
	return 0
}

func (e *BinaryExpr) String() string {
	return fmt.Sprintf("(%v %s %v)", e.a, string(e.op), e.b)
}

func (e *UnaryExpr) String() string {
	return fmt.Sprintf("%d", e.a)
}

type TokenType int

const (
	TokenTypeNumber = iota
	TokenTypeOp
	TokenTypeParenOpen
	TokenTypeParenClose
)

type Token struct {
	Type      TokenType
	ValueRune rune
	ValueInt  int
}

type TokenScanner struct {
	Tokens []Token
	i      int
}

func (s *TokenScanner) Scan() bool {
	return s.i < len(s.Tokens)
}

func (s *TokenScanner) Peek() Token {
	return s.Tokens[s.i]
}

func (s *TokenScanner) Next() Token {
	t := s.Tokens[s.i]
	s.i++
	return t
}

func lexNumber(s *scanner.Scanner) int {
	nums := ""
	for unicode.IsDigit(s.Peek()) {
		nums += string(s.Next())
	}
	num, err := strconv.Atoi(nums)
	if err != nil {
		panic(err)
	}
	return num
}

func lex(line string) []Token {
	tokens := make([]Token, 0)
	s := scanner.Scanner{}
	s.Init(strings.NewReader(line))
	for {
		switch s.Peek() {
		case ' ':
			s.Next()
		case '(':
			tokens = append(tokens, Token{
				Type: TokenTypeParenOpen,
			})
			s.Next()
		case ')':
			tokens = append(tokens, Token{
				Type: TokenTypeParenClose,
			})
			s.Next()
		case '+', '*':
			tokens = append(tokens, Token{
				Type:      TokenTypeOp,
				ValueRune: s.Next(),
			})
		default:
			r := s.Peek()
			if unicode.IsDigit(r) {
				tokens = append(tokens, Token{
					Type:     TokenTypeNumber,
					ValueInt: lexNumber(&s),
				})
				break
			}
			if r < 0 {
				return tokens
			}
			panic("invalid input")
		}
	}
}

func parseGroup(s *TokenScanner) Expr {
	var expr *BinaryExpr
Loop:
	for s.Scan() {
		switch s.Peek().Type {
		case TokenTypeNumber:
			e := &UnaryExpr{
				a: s.Next().ValueInt,
			}
			if expr == nil {
				expr = &BinaryExpr{
					a: e,
				}
			} else {
				expr.b = e
				expr = &BinaryExpr{
					a: expr,
				}
			}
		case TokenTypeOp:
			expr.op = s.Next().ValueRune
			if s.Peek().Type == TokenTypeNumber {
				expr.b = &UnaryExpr{
					a: s.Next().ValueInt,
				}
				expr = &BinaryExpr{
					a: expr,
				}
			}
		case TokenTypeParenOpen:
			s.Next()
			e := parseGroup(s)
			if expr == nil {
				expr = &BinaryExpr{
					a: e,
				}
			} else {
				expr.b = e
				expr = &BinaryExpr{
					a: expr,
				}
			}
		case TokenTypeParenClose:
			s.Next()
			break Loop
		}
	}
	if expr.b == nil {
		return expr.a
	}
	return expr
}

func parse(tokens []Token) Expr {
	s := TokenScanner{
		Tokens: tokens,
	}
	return parseGroup(&s)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	sum := 0
	for s.Scan() {
		tokens := lex(s.Text())
		expr := parse(tokens)
		res := expr.Eval()
		sum += res
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", sum)
}
