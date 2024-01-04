package dunsinane

import (
	"reflect"
	"testing"
)

func TestTextProcessing(t *testing.T) {
	text := "[line The [bold quick bold] brown fox line] [line jumped [bold over bold] the lazy dog. line]"
	index := ProcessDocument(text)

	checkIndex(t, index, "dog", []Extent{
		{start: 8, end: 8},
	})

	checkIndex(t, index, "[bold", []Extent{
		{start: 0.5, end: 0.5},
		{start: 4.5, end: 4.5},
	})
}

func checkIndex(t *testing.T, index Index, token Token, expected []Extent) {
	actualList := index[token].extents

	if !reflect.DeepEqual(actualList, expected) {
		t.Fatalf(`index[%q] = %v, wanted %v`, token, actualList, expected)
	}
}
