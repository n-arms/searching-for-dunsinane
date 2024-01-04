package dunsinane

import (
	"strings"
	"unicode"
)

type Token = string
type Position = float32

type Index = map[Token]GCList

type GCList struct {
	extents []Extent
}

type Extent struct {
	start Position
	end   Position
}

func isMarkup(token Token) bool {
	return strings.HasPrefix(token, "[") || strings.HasSuffix(token, "]")
}

func tokenize(document string) []Token {
	tokens := []Token{}
	current := strings.Builder{}

	for _, char := range document {
		if unicode.IsPunct(char) {
			continue
		} else if unicode.IsSpace(char) {
			tokens = append(tokens, current.String())
			current = strings.Builder{}
		} else {
			current.WriteRune(char)
		}
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}

const EPSILON Position = 0.5

func index(tokens []Token) Index {
	index := Index{}

	nextIndex := float32(0)

	for _, token := range tokens {
		var documentIndex Position
		if isMarkup(token) {
			documentIndex = nextIndex - EPSILON
		} else {
			documentIndex = nextIndex
		}
		extent := Extent{
			start: documentIndex,
			end:   documentIndex,
		}
		list, exists := index[token]
		if exists {
			list.extents = append(list.extents, extent)
			index[token] = list
		} else {
			index[token] = GCList{
				extents: []Extent{extent},
			}
		}
	}

	return index
}

func ProcessDocument(document string) Index {
	return index(tokenize(document))
}
