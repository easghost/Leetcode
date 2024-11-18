func searchRange(nums []int, target int) []int {
	first, last := -1, -1

	for i, v := range nums {
		if v == target {
			if first == -1 {
				first = i
			}

			last = i
		}
	}

	return []int{first, last}
}