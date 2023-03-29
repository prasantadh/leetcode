func isAlphaNum(b byte) bool {
    if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
        return true
    }
    return false
}
func isPalindrome(s string) bool {

    i, j := 0, len(s) - 1
    s = strings.ToLower(s)

    for i < j {
        if !isAlphaNum(s[i]) {
            i++
        } else if !isAlphaNum(s[j]) {
            j--
        } else if s[i] != s[j] {
            return false
        } else {
            i++
            j--
        }
    }
    return true
    
}
