class Solution {
    fun checkInclusion(s1: String, s2: String): Boolean {
	val s1Arr = IntArray(26)
	val s2Arr = IntArray(26)
	
	for (v in s1) {
		s1Arr[v - 'a']++
	}
	
	for (i in s2.indices) {
		val v = s2[i]
		if (i >= s1.length) {
			s2Arr[s2[i - s1.length] - 'a']--
		}
		s2Arr[v - 'a']++
		if (s1Arr.contentEquals(s2Arr)) {
			return true
		}
	}
	
	return false
}
}