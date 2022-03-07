class Solution {
    fun sortedSquares(A: IntArray): IntArray {
        var leftMarker = 0
        var rightMarker = A.size - 1

        var resultIndex = A.size - 1
        val result = IntArray(A.size)

        while (leftMarker <= rightMarker) {
            val left = Math.abs(A[leftMarker])
            val right = Math.abs(A[rightMarker])

            result[resultIndex] = if (right > left) {
                rightMarker--
                right * right
            } else {
                leftMarker++
                left * left
            }

            resultIndex--
        }

        return result
    }
}