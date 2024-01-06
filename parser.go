package dunsinane

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	variableKind    = iota
	stringKind      = iota
	bothOfKind      = iota
	containsKind    = iota
	containedInKind = iota
	orderKind       = iota
	eofKind         = iota
)

func (token QToken) isAtom() bool {
	return token.kind == stringKind || token.kind == variableKind
}

func (token QToken) isEof() bool {
	return token.kind == eofKind
}

func (token QToken) isOperator() bool {
	return token.kind == bothOfKind || token.kind == containsKind || token.kind == containedInKind || token.kind == orderKind
}

type TokenKind = int

type QToken struct {
	kind TokenKind
	text string
}

type Lexer struct {
	tokens []QToken
	index  int
}

func makeLexer(input string) Lexer {
	current := strings.Builder{}
	inToken := false
	isString := false

	tokens := []QToken{}

	for _, char := range input {
		if inToken {
			if isString {
				if char == '"' {
					tokens = append(tokens, QToken{kind: stringKind, text: current.String()})
					inToken = false
					current = strings.Builder{}
				} else {
					current.WriteRune(char)
				}
			} else {
				if unicode.IsDigit(char) || unicode.IsLetter(char) {
					current.WriteRune(char)
				} else {
					tokens = append(tokens, QToken{kind: variableKind, text: current.String()})
					inToken = false
					current = strings.Builder{}
				}
			}
		} else {
			switch char {
			case '"':
				isString = true
				inToken = true
			case '^':
				tokens = append(tokens, QToken{kind: bothOfKind, text: string(char)})
			case '>':
				tokens = append(tokens, QToken{kind: containsKind, text: string(char)})
			case '<':
				tokens = append(tokens, QToken{kind: containedInKind, text: string(char)})
			case '*':
				tokens = append(tokens, QToken{kind: orderKind, text: string(char)})
			default:
				if !unicode.IsSpace(char) {
					isString = false
					inToken = true
					current.WriteRune(char)
				}
			}
		}
	}

	fmt.Println("Lexed", input, "into tokens", tokens)

	return Lexer{
		tokens: tokens,
		index:  0,
	}
}

func (l *Lexer) next() QToken {
	if l.index < len(l.tokens) {
		l.index += 1
		return l.tokens[l.index-1]
	} else {
		return QToken{kind: eofKind, text: ""}
	}
}

func (l *Lexer) peek() QToken {
	if l.index < len(l.tokens) {
		return l.tokens[l.index]
	} else {
		return QToken{kind: eofKind, text: ""}
	}
}

func ParseExpr(input string, index Index) (Query, error) {
	l := makeLexer(input)
	return exprWithPower(&l, 0, make(map[string]Query), index)
}

func infixBindingPower(operator QToken) (int, int) {
	kind := operator.kind

	if kind == bothOfKind {
		return 1, 2
	} else if kind == containedInKind || kind == containsKind {
		return 3, 4
	} else {
		return 5, 6
	}
}

func exprWithPower(lexer *Lexer, minPower int, variables map[string]Query, index Index) (Query, error) {
	fmt.Println("recur with tokens", lexer)
	lhsToken := lexer.next()
	if !lhsToken.isAtom() {
		return nil, errors.New(fmt.Sprintf("expected atom, got %v", lhsToken))
	}
	var result Query
	var exists bool

	if lhsToken.kind == stringKind {
		result, exists = index[lhsToken.text]
	} else {
		result, exists = variables[lhsToken.text]
	}

	if !exists {
		return nil, errors.New(fmt.Sprintf("unknown atom %v", lhsToken.text))
	}

	fmt.Println("parsed lhs", result)

	for {
		op := lexer.peek()
		if op.isEof() {
			break
		} else if op.isOperator() {

		} else {
			return nil, errors.New(fmt.Sprintf("expected operator, got %v", op))
		}
		leftPower, rightPower := infixBindingPower(op)

		if leftPower < minPower {
			break
		}

		lexer.next()
		rhs, err := exprWithPower(lexer, rightPower, variables, index)
		if err != nil {
			return nil, err
		}
		if op.kind == bothOfKind {
			result = BothOf{
				first:  result,
				second: rhs,
			}
		} else if op.kind == containsKind {
			result = Contains{
				first:  result,
				second: rhs,
			}
		} else if op.kind == containedInKind {
			result = ContainedIn{
				first:  result,
				second: rhs,
			}
		} else if op.kind == orderKind {
			result = Order{
				first:  result,
				second: rhs,
			}
		}
	}

	return result, nil
}
