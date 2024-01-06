package dunsinane

import "testing"

func TryParse(t *testing.T, query string, index Index, expected Query) {
	actual, err := ParseExpr(query, index)
	if err != nil {
		t.Fatalf("Parse error %v while parsing %s into %v\n", err, query, actual)
	}
	CheckEqual(t, actual, expected)

}

func TestParser(t *testing.T) {
	index := Index{
		"x": GCList{extents: []Extent{{start: 0, end: 0}}},
		"y": GCList{extents: []Extent{{start: 1, end: 1}}},
	}

	TryParse(t, `"x"`, index, index["x"])
	TryParse(t, `"x" * "y"`, index, Order{first: index["x"], second: index["y"]})
	TryParse(t, `"x" ^ "y"`, index, BothOf{first: index["x"], second: index["y"]})
	TryParse(t, `"x" > "y"`, index, Contains{first: index["x"], second: index["y"]})
	TryParse(t, `"x" < "y"`, index, ContainedIn{first: index["x"], second: index["y"]})
	TryParse(t, `"x" * "y" > "y"`, index, Contains{first: Order{first: index["x"], second: index["y"]}, second: index["y"]})
	TryParse(t, `"x" < "x" * "y"`, index, ContainedIn{first: index["x"], second: Order{first: index["x"], second: index["y"]}})
}
