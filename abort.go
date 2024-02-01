package piscine

func Abort(a, b, c, d, e int) int {
	num := 0
	arr := []int{a, b, c, d, e}
	var arr1 []int
	arr1 = sort(arr)
	for i := 0; i < len(arr1); i++ {
		num = arr1[len(arr1)/2]
	}
	return num
}

func Sort1(arr1 []int) []int {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1)-i-1; j++ {
			if arr1[j] > arr1[j+1] {
				t := arr1[j]
				arr1[j] = arr1[j+1]
				arr1[j+1] = t
			}
		}
	}
	return arr1
}
