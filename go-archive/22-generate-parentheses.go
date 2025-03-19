func generateParenthesis(n int) []string {

    /****
    backtracking problem
    for each open parenthesis, there must be a close one.
    keep track of how many open and close parenthesis you have left
    use them
    *******/
    ans := []string{}
    temp := []byte{}
    var helper func(open int, close int)
    helper = func (open int, close int) {
        
        if open > close {
            return
        }

        if open == 0 && close == 0 {
            ans = append(ans, string(temp))
            return
        }

        if open > 0 {
            temp = append(temp, '(')
            helper(open - 1, close)
            temp = temp[:len(temp) - 1]
        }

        if close > 0 {
            temp = append(temp, ')')
            helper(open, close - 1)
            temp = temp[:len(temp) - 1]
        }


    }

    helper(n, n)
    return ans
    
}
