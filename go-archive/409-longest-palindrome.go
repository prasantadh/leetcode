func longestPalindrome(s string) int {
    /* strategy: get the even occurences of each character,
    at the end get one more if we didn't get all the characters */
    memo := [128]int{}
    for i := 0; i < len(s); i++ {
        memo[s[i]] += 1
    }
    
    ans := 0
    for i := 0; i < len(memo); i++ {
        ans += memo[i]
        if memo[i] & 1 == 1 {
            ans -= 1
        }
    }
    if len(s) > ans {
        ans++
    }
    return ans
    /* because I didn't read the question properly at first
    I also solved for longest palindrome that exists in a given string
    */

    // memo := [2001][2001]bool{}

    // n := len(s)
    // memo[0][0] = true
    // for r := 1; r < n; r++ {
    //     memo[r][r] = true
    //     memo[r][r - 1] = true
    // }

    // ans := 1
    // for r := n - 2; r >= 0; r-- {
    //     for c := r + 1; c < n; c++ {
    //         memo[r][c] = memo[r + 1][c - 1] && s[r] == s[c]
    //         if t := c - r + 1; memo[r][c] && t > ans {
    //             fmt.Println(r, c)
    //             ans = t
    //         }
    //     }
    // } 
    // return ans
}
