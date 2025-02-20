package utils

func RemoveElement[T any](slice *[]T, index int) {
	if slice == nil {
		panic("nil slice")
	}
	if index < 0 || index >= len(*slice) {
		panic("Remove Element: Index out of range")
	}
	// Delete index
	aux := *slice
	*slice = append(aux[:index], aux[index+1:]...)
}

func Swap[T any](slice *[]T, i, j int) {
	aux := *slice
	if i >= 0 && i < len(aux) && j >= 0 && j < len(aux) {
		aux[i], aux[j] = aux[j], aux[i]
	}
	*slice = aux
}

func IsStraight(x int) bool {
	valids := []int{0, 4, 7, 8, 9}
	//We keep everything true while testing
	isValid := true
	for _, v := range valids {
		if v == x {
			isValid = true
		}
	}
	return isValid

}

func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
