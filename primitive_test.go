package dunsinane

import (
	"reflect"
	"testing"
)

func TestGCListAccess(t *testing.T) {
	var query Query = ExampleIndex()["[line"]

	CheckEqual(t, query.accessT(3), Extent{start: 3.5, end: 3.5})
	CheckEqual(t, query.accessP(3), Extent{start: 3.5, end: 3.5})
	CheckEqual(t, query.accessTPrime(3), Extent{start: -0.5, end: -0.5})
	CheckEqual(t, query.accessPPrime(3), Extent{start: -0.5, end: -0.5})
}

func CheckEqual[T any](t *testing.T, actual T, wanted T) {
	if !reflect.DeepEqual(actual, wanted) {
		//debug.PrintStack()
		t.Fatalf("got %v, wanted %v", actual, wanted)
	}
}

func TestSigmaAccess(t *testing.T) {
	var query Query = Sigma{
		length: EPSILON + 1,
	}

	CheckEqual(t, query.accessT(2), Extent{start: 2, end: 2 + 1 + EPSILON - EPSILON})
	CheckEqual(t, query.accessP(2), Extent{start: 2 - 1 - EPSILON + EPSILON, end: 2})
}
