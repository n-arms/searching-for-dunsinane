package dunsinane

import (
	"fmt"
	"math"
	"reflect"
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
	CheckEqual(t, query.accessP(6).isReal(), false)
	CheckEqual(t, query.accessTPrime(0).isReal(), false)
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
	CheckEqual(t, query.accessPPrime(4).isReal(), false)
}

func TestContainedIn(t *testing.T) {
	index := ExampleIndex()

	var bold Query = Order{
		first:  index["[bold"],
		second: index["bold]"],
	}

	var query Query = ContainedIn{
		first:  index["over"],
		second: bold,
	}

	CheckEqual(t, query.accessP(0), Extent{start: 5, end: 5})
	CheckEqual(t, query.accessTPrime(6), Extent{start: 5, end: 5})
	CheckEqual(t, query.accessPPrime(4).isReal(), false)
}

func TestBothOf(t *testing.T) {
	index := ExampleIndex()

	var query Query = BothOf{
		first:  index["quick"],
		second: index["brown"],
	}

	CheckEqual(t, query.accessT(0), Extent{start: 1, end: 2})
	CheckEqual(t, query.accessPPrime(3), Extent{start: 1, end: 2})
}

func drive(query Query) {
	k := NegInfinity()

	for {
		e := query.accessT(k)
		if e.isReal() {
			k = e.end
			fmt.Println(e)
		} else {
			fmt.Printf("Finished with invalid extent %v\n", e)
			break
		}
	}
}

func drivePPrime(query Query) {
	k := Infinity()
	var prev *Extent = nil

	for {
		e := query.accessPPrime(k)
		if e.isReal() {
			if prev != nil && reflect.DeepEqual(*prev, e) {
				k -= EPSILON
				continue
			} else {
				fmt.Println(e)
				prev = new(Extent)
				*prev = e
				if math.IsInf(float64(k), 0) {
					k = e.start - EPSILON
				}
			}
		} else {
			fmt.Printf("Finished with invalid extent %v\n", e)
			break
		}
	}
}
