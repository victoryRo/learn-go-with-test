package array

func Sum(array []int) (sum int) {

	for _, value := range array {
		sum += value
	}

	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, value := range numbersToSum {
		if len(value) == 0 {
			sums = append(sums, 0)
		} else {
			tail := value[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
