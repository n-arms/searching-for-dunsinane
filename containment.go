package dunsinane

// A ▷ B in the original paper, the extents in `first` that contain any extents in `second`
type Contains struct {
	first  Query
	second Query
}

func (c Contains) accessT(k Position) Extent {
	e := c.first.accessT(k)
	return c.accessP(e.end)
}

func (c Contains) accessP(k Position) Extent {
	// find first that ends at or after k
	e := c.first.accessP(k)
	// find first that starts inside e
	ePrime := c.second.accessT(e.start)

	// check if they nest
	if ePrime.end <= e.end {
		return e
	} else {
		return c.accessP(ePrime.end)
	}
}

func (c Contains) accessTPrime(k Position) Extent {
	e := c.first.accessTPrime(k)
	return c.accessPPrime(e.start)
}

// find last that starts at or before k
func (c Contains) accessPPrime(k Position) Extent {
	e := c.first.accessPPrime(k)
	ePrime := c.second.accessTPrime(e.end)

	if ePrime.start >= e.start {
		return e
	} else {
		return c.accessPPrime(ePrime.start)
	}
}

// A ◁ B in the original paper, the extents in `first` that are contained in any extents in `second`
type ContainedIn struct {
	first  Query
	second Query
}

func (c ContainedIn) accessT(k Position) Extent {
	e := c.second.accessT(k)
	return c.accessP(e.end)
}

func (c ContainedIn) accessP(k Position) Extent {
	e := c.second.accessP(k)
	ePrime := c.first.accessT(e.start)

	if ePrime.end <= e.end {
		return ePrime
	} else {
		return c.accessP(ePrime.end)
	}
}

func (c ContainedIn) accessTPrime(k Position) Extent {
	e := c.second.accessTPrime(k)
	return c.accessPPrime(e.start)
}

func (c ContainedIn) accessPPrime(k Position) Extent {
	e := c.second.accessPPrime(k)
	ePrime := c.first.accessTPrime(e.end)

	if ePrime.start >= e.start {
		return ePrime
	} else {
		return c.accessPPrime(ePrime.start)
	}
}
