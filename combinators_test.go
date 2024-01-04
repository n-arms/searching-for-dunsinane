package dunsinane

import (
	"fmt"
	"math"
	"testing"
)

func TestOrder(t *testing.T) {
	index := ExampleIndex()

	var query Query = Order{
		first:  index["[bold"],
		second: index["bold]"],
	}

	CheckEqual(t, query.accessT(0), Extent{start: 0.5, end: 1.5})
	CheckEqual(t, query.accessT(1), Extent{start: 4.5, end: 5.5})
	CheckEqual(t, query.accessTPrime(1000), Extent{start: 4.5, end: 5.5})
	CheckEqual(t, query.accessP(0), Extent{start: 0.5, end: 1.5})
}

func TestContains(t *testing.T) {
	index := ExampleIndex()

	var bold Query = Order{
		first:  index["[bold"],
		second: index["bold]"],
	}

	var query Query = Contains{
		first:  bold,
		second: index["over"],
	}

	CheckEqual(t, query.accessP(0), Extent{start: 4.5, end: 5.5})
}

func drive(query Query, index Index) {
	k := float32(math.Inf(-1))
	for k != float32(math.Inf(1)) {
		e := query.accessT(k)
		fmt.Println(e)
		k = e.end
	}
}
