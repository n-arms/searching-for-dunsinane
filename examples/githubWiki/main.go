package main

import (
	"fmt"
	"strings"
)

func main() {
	_, tokens := ProcessWiki("https://github.com/mjansen4857/pathplanner.wiki.git")
	document := strings.Builder{}

	for _, token := range tokens {
		document.WriteString(token)
		document.WriteRune(' ')
	}

	fmt.Println(document.String())
}
