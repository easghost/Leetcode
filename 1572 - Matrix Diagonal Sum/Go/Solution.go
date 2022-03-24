func diagonalSum(mat [][]int) int {
    sum := 0
    rows := len(mat)
    if rows == 0{
        return sum
    }
    for i := 0; i < rows; i++ {
        sum += mat[i][i]
        sum += mat[i][rows-i-1]
    }
    if rows & 1 == 1 {
        sum -= mat[rows>>1][rows>>1]
    }
    
    return sum
}