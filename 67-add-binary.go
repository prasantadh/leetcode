func addBinary(a string, b string) string {
  /* start at the back of the string,
  add one digit at a time */
    
    if len(b) > len(a) {
        a, b = b, a
    }

    ans := make([]byte, len(a))

    carry := false
    for i := 1; i <= len(a); i++ {
        ai := a[len(a) - i] == '1'
        bi := false
        if i <= len(b) {
            bi = b[len(b) - i] == '1'
        }
        sum := (ai != bi != carry) 
        carry = (ai && bi) || (ai && carry) || (bi && carry)
        if sum {
            ans[len(a) - i] = '1'
        } else {
            ans[len(a) - i] = '0'
        }
    }
    if carry {
        return string(append([]byte{'1'}, ans...))
    }
    return string(ans)
    
    
}
