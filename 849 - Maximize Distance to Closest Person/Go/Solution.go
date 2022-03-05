func maxDistToClosest(seats []int) int {
	l := len(seats)
	prev := -1
	ans := 0

	for i := 0; i < l; i++ {
		if seats[i] == 1 {
			if prev == -1 {
				ans = i
			} else {
				ans = max(ans, (i-prev)/2)
			}
			prev = i
		}
	}
	ans = max(ans, l-prev-1)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}