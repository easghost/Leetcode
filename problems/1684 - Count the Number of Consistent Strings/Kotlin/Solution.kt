class Solution {
    fun countConsistentStrings(allowed: String, words: Array<String>): Int {
        var count = words.size
        val wordSet = IntArray(26)

        for (i in allowed.indices) {
            wordSet[allowed[i] - 'a']++
        }

        for (word in words) {
            for (i in word.indices) {
                if (wordSet[word[i] - 'a'] == 0) {
                    count--
                    break
                }
            }
        }

        return count
    }
}