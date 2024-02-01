package piscine

func Unmatch(a []int) int {
	f := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				f++
			}
		}
		if f%2 != 0 {
			return a[i]
		}
	}
	return -1
}
