package dunsinane

func emptyExtent() Extent {
	return Extent{
		start: Infinity(),
		end:   Infinity(),
	}
}

func negEmptyExtent() Extent {
	return Extent{
		start: NegInfinity(),
		end:   NegInfinity(),
	}
}

func (l GCList) accessT(k Position) Extent {
	for _, extent := range l.extents {
		if extent.start >= k {
			return extent
		}
	}
	return emptyExtent()
}

func (l GCList) accessP(k Position) Extent {
	for _, extent := range l.extents {
		if extent.end >= k {
			return extent
		}
	}
	return emptyExtent()
}

func (l GCList) accessTPrime(k Position) Extent {
	last := len(l.extents) - 1
	for i := range l.extents {
		extent := l.extents[last-i]
		if extent.end <= k {
			return extent
		}
	}
	return negEmptyExtent()
}

func (l GCList) accessPPrime(k Position) Extent {
	last := len(l.extents) - 1
	for i := range l.extents {
		extent := l.extents[last-i]
		if extent.start <= k {
			return extent
		}
	}
	return negEmptyExtent()
}

// All extents of the given length
type Sigma struct {
	length Position
}

func (s Sigma) accessT(k Position) Extent {
	return Extent{
		start: k,
		end:   k + s.length - EPSILON,
	}
}

func (s Sigma) accessP(k Position) Extent {
	return Extent{
		start: k - s.length + EPSILON,
		end:   k,
	}
}

func (s Sigma) accessTPrime(k Position) Extent {
	return s.accessT(k)
}

func (s Sigma) accessPPrime(k Position) Extent {
	return s.accessP(k)
}
