func longestPalindrome(s string) string {
  /* code copy pasted from 409 longest palindrome.
  consult the solution there for an explanation */
    memo := [1001][1001]bool{}

    n := len(s)
    memo[0][0] = true
    for r := 1; r < n; r++ {
        memo[r][r] = true
        memo[r][r - 1] = true
    }

    ans := s[:1]
    anslen := 1
    for r := n - 2; r >= 0; r-- {
        for c := r + 1; c < n; c++ {
            memo[r][c] = memo[r + 1][c - 1] && s[r] == s[c]
            if t := c - r + 1; memo[r][c] && t > anslen {
                ans = s[r:c+1]
                anslen = t
            }
        }
    } 
    return ans
}
