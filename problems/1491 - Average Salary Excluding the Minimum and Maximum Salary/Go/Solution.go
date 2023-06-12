func average(salary []int) float64 {
    sum, min, max := 0, 1000000, 1000
    for _, v := range salary {
        sum += v
        if v < min { min = v }
        if v > max { max = v }
    }
    return float64(sum - min - max) / float64(len(salary) - 2)
}