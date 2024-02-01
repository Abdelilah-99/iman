package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	m := make(map[string]int)
	count := 0
	arr := Split(str, " ")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		m[arr[i]] = count
		count = 0
	}
	return m
}
