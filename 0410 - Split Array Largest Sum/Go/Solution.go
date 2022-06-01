func splitArray(nums []int, m int) int {
    var low, high int = 0, 0
    for _, v := range nums {
        if low < v { low = v }
        high += v
    }

    for low <= high {
        mid := (low + high) / 2
        res := canPartition(nums, mid)
        if res <= m {
            high = mid - 1 
        } else {
            low = mid + 1
        }
    }
    return low
}

func canPartition(n []int, goal int) int {
    var cur, p int
    for _, v := range n {
        cur += v
        if cur > goal {
            cur = v
            p++
        }
    }
    return p + 1
}