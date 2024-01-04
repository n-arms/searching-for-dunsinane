package dunsinane

// A ◇ B in the original paper, the extents in first merged with all the extents in second that follow them
type Order struct {
	first  Query
	second Query
}

func (o Order) accessT(k Position) Extent {
	// find the first extent in A at or after K
	e := o.first.accessT(k)
	// find the first extent in B after (non overlapping) e
	ePrime := o.second.accessT(e.end + EPSILON)
	// find the closest extent in A that comes before ePrime
	eDoublePrime := o.first.accessTPrime(ePrime.start - EPSILON)

	if eDoublePrime.start == Infinity() || ePrime.end == Infinity() {
		return emptyExtent()
	} else {
		return Extent{
			start: eDoublePrime.start,
			end:   ePrime.end,
		}
	}
}

func (o Order) accessP(k Position) Extent {
	e := o.accessTPrime(k - EPSILON)
	return o.accessT(e.start + EPSILON)
}

// find the last that starts at or before k
func (o Order) accessPPrime(k Position) Extent {
	e := o.accessT(k + EPSILON)
	return o.accessTPrime(e.end - EPSILON)
}

// find the last that ends at or before
func (o Order) accessTPrime(k Position) Extent {
	e := o.second.accessTPrime(k)
	ePrime := o.first.accessTPrime(e.start - EPSILON)
	eDoublePrime := o.second.accessT(ePrime.end + EPSILON)

	return Extent{
		start: ePrime.start,
		end:   eDoublePrime.end,
	}
}
