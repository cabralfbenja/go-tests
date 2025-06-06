package integers

func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func SumAll(args ...[]int) []int {
	var sums []int
	for _, array := range args {
		sums = append(sums, Sum(array))
	}
	return sums
}

func SumAllTails(args ...[]int) []int {
	var sums []int
	for _, array := range args {
        if len(array) == 0 {
            sums = append(sums, 0)
            continue
        }
		tail := array[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
