func nextGreaterElement(nums1 []int, nums2 []int) []int {
    tmp := make(map[int]int)
    for i := 0; i < len(nums2); i++ {
        d := - 1
        for j := i + 1; j < len(nums2); j++ {
            if nums2[j] > nums2[i] { d = nums2[j]; break }
        }
        tmp[nums2[i]] = d
    }
    ans := []int{}
    for _, n := range nums1 {
        ans = append(ans, tmp[n])
    }
    return ans
}