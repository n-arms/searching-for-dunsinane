package dunsinane

import (
	"reflect"
	"testing"
)

func TestGCListAccess(t *testing.T) {
	list := ExampleIndex()["[line"]

	checkEqual(t, list.accessT(3), Extent{start: 3.5, end: 3.5})
	checkEqual(t, list.accessP(3), Extent{start: 3.5, end: 3.5})
	checkEqual(t, list.accessTPrime(3), Extent{start: -0.5, end: -0.5})
	checkEqual(t, list.accessPPrime(3), Extent{start: -0.5, end: -0.5})
}

func checkEqual[T any](t *testing.T, actual T, wanted T) {
	if !reflect.DeepEqual(actual, wanted) {
		t.Fatalf("got %v, wanted %v", actual, wanted)
	}
}
