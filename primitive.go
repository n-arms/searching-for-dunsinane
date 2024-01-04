package dunsinane

import "math"

func emptyExtent() Extent {
	return Extent{
		start: float32(math.Inf(+1)),
		end:   float32(math.Inf(+1)),
	}
}

// TODO: reimplement list queries using binary search

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
	return emptyExtent()
}

func (l GCList) accessPPrime(k Position) Extent {
	last := len(l.extents) - 1
	for i := range l.extents {
		extent := l.extents[last-i]
		if extent.start <= k {
			return extent
		}
	}
	return emptyExtent()
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
