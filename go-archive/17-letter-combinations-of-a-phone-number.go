func letterCombinations(digits string) []string {
    
    if len(digits) == 0 {
        return []string{}
    }

    var dialpad [][]byte = [][]byte {
        {},
        {},
        {'a','b','c'},
        {'d','e','f'},
        {'g','h','i'},
        {'j','k','l'},
        {'m','n','o'},
        {'p','q','r','s'},
        {'t','u','v'},
        {'w','x','y','z'},
    }

    ans := []string{}
    temp := []byte{}
    var bt func(i int)
    bt = func(i int) {
        if i == len(digits) {
            ans = append(ans, string(temp))
            return
        }

        for _, val := range(dialpad[digits[i] - '0']){
            // fmt.Println(digits[i] - '0')
            temp = append(temp, val)
            bt(i+1)
            temp = temp[:len(temp) - 1]
        }
    }
    bt(0)
    return ans

}
