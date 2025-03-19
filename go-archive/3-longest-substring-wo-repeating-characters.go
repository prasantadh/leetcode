func lengthOfLongestSubstring(s string) int {
    // head := the current character
    // tail := the last character such that there is no
    //          repetition between (tail, head]
    // iterate over all heads, take the maximum of head - tail
    
    n := len(s)
    if n == 0 {
        return n
    }

    lastseen := [128]int{}
    for i := range lastseen {
        lastseen[i] = -1
    }

    tail := -1
    head := 0
    ans  := 0
    for head < n {
        c := s[head]
        i := lastseen[c]
        if i > tail {
            tail = i
        }
        if t := head - tail; t > ans {
            ans = t
        }
        lastseen[c] = head
        head++
    }
    return ans
}
