func generate(numRows int) [][]int {
    var ans [][]int
    for i := 0; i < numRows; i++ {
        ans = append(ans, make([]int, i + 1))
        current := ans[len(ans) - 1]
        current[0] = 1
        current[len(current) - 1] = 1
        for j := 1; j < len(current) - 1; j++ {
            current[j] = ans[i-1][j] + ans[i-1][j-1]
        }
    }
    return ans   
}
