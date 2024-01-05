package dunsinane

type BothOf struct {
	first  Query
	second Query
}

// A △ B in the original paper, extents that contain extents from both A and B
func (b BothOf) accessT(k Position) Extent {
	first := b.first.accessT(k)
	second := b.second.accessT(k)
	firstPrime := b.first.accessTPrime(max(first.end, second.end))
	secondPrime := b.second.accessTPrime(max(first.end, second.end))
	return Extent{
		start: min(firstPrime.start, secondPrime.start),
		end:   max(firstPrime.end, secondPrime.end),
	}
}

func (b BothOf) accessP(k Position) Extent {
	e := b.accessTPrime(k - EPSILON)
	return b.accessT(e.start + EPSILON)
}

func (b BothOf) accessTPrime(k Position) Extent {
	first := b.first.accessTPrime(k)
	second := b.second.accessTPrime(k)
	firstPrime := b.first.accessT(min(first.start, second.start))
	secondPrime := b.second.accessT(min(first.start, second.start))
	return Extent{
		start: min(firstPrime.start, secondPrime.start),
		end:   max(firstPrime.end, secondPrime.end),
	}
}

func (b BothOf) accessPPrime(k Position) Extent {
	e := b.accessT(k + EPSILON)
	return b.accessTPrime(e.end - EPSILON)
}
