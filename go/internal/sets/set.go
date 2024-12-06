package sets

type IntSet = map[int]bool

func Insert(val int, set IntSet) IntSet {
	if set == nil {
		set = make(IntSet)
	}
	set[val] = true
	return set
}
