package offer

func minArray11(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		mid := (i + j) / 2
		if numbers[mid] > numbers[j] {
			i = mid + 1
		} else if numbers[mid] < numbers[j] {
			j = mid
		} else {
			j--
		}
	}
	return numbers[i]
}
