package code

func Execuate_167(numbers []int, target int) []int {
	return twoSum(numbers, target)
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	sum := 0
	for i < j {
		sum = numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}

	}

	return []int{0, 0}

}
