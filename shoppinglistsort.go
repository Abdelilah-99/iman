package piscine

func ShoppingListSort(slice []string) []string {
	var str []string
	str = sort1(slice)
	return str
}

func sort1(arr1 []string) []string {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1)-i-1; j++ {
			if len(arr1[j]) > len(arr1[j+1]) {
				t := arr1[j]
				arr1[j] = arr1[j+1]
				arr1[j+1] = t
			}
		}
	}
	return arr1
}
