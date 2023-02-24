// Runtime: 16 ms
// Memory Usage: 3.5 MB

func smallerNumbersThanCurrent(nums []int) []int {
	var res []int

	for i, val1 := range nums {
		counter := 0
		for j, val2 := range nums {
			if i != j && val2 < val1 {
				counter++
			}
		}
		res = append(res, counter)
	}

	return res
}