import "math"

func myAtoi(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }

    ans := 0
    i := 0

    for i < n && s[i] == ' ' {
        i++
    }
    if i == n {
        return 0
    }

    isNeg := s[i] == '-'
    if isNeg || s[i] == '+' {
        i++
    }

    for i < n && s[i] == '0' {
        i++
    }
    if i == n {
        return 0
    }

    ref := i
    for i < n && s[i] >= '0' && s[i] <= '9' {
        ans = ans * 10 + int(s[i]) - '0'
        i++
        if i - ref + 1 > 11 {
            break
        }
    }
    
    if isNeg {
        ans = -ans
        if ans < math.MinInt32 {
            return math.MinInt32
        }
    }

    if ans > math.MaxInt32 {
        return math.MaxInt32
    }
    return ans; 
}
