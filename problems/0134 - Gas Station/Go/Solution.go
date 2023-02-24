// Runtime: 52 ms
// Memory Usage: 10.1 MB

func canCompleteCircuit(gas []int, cost []int) int {
	sumGas := 0
	sumCost := 0
	for _, g := range gas {
		sumGas += g
	}

	for _, c := range cost {
		sumCost += c
	}

	if sumGas < sumCost {
		return -1
	}

	sum := 0
	start := 0
	next := 0
	for {
		diff := gas[next] - cost[next]
		sum = sum + diff
		next = (next + 1) % len(gas)
		if sum >= 0 {
			if start == next {
				return start
			}
			continue
		}

		start = next
		sum = 0
	}

	return -1
}