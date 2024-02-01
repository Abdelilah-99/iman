package piscine

func MakeRange(min, max int) []int {
	if min < max && (min != 0 || max != 0) {
		myslice := make([]int, max-min)
		if min < max {
			j := min
			for i := min; i < max; i++ {
				myslice[i-j] = i
			}
		}
		return myslice
	}
	return []int(nil)
}
