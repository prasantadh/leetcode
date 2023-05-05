func maxVowels(s string, k int) int {

    temp := 0
    for i := 0; i < k; i++ {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
            temp += 1
        }
    }

    ans := temp
    for i := k; i < len(s); i++ {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
            temp += 1
        }
        if l := s[i-k]; l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' {
            temp -= 1
        }
        if temp > ans {
            ans = temp
        }
    }
    return ans
}
