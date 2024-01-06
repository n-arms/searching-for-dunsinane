package dunsinane

import "fmt"

type Query interface {
	// Find the first extent starting at or after k
	accessT(k Position) Extent
	// Find the first extent ending at or after k
	accessP(k Position) Extent
	// Find the last extent ending at or before k
	accessTPrime(k Position) Extent
	// Find the last extent starting at or before k
	accessPPrime(k Position) Extent
}

func DriveQuery(query Query) {
	k := NegInfinity()

	for {
		e := query.accessT(k)

		if e.isReal() {
			fmt.Println(e)
		} else {
			break
		}

		k = e.end + EPSILON
	}
}
