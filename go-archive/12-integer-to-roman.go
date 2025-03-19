type numeral struct {
    value int
    symbol []byte
}

func intToRoman(num int) string {
      
    numerals := []numeral{
        {value: 1000, symbol: []byte{'M'} },
        {900, []byte{'C', 'M'}},
        {500, []byte{'D'}},
        {400, []byte{'C', 'D'}},
        {100, []byte{'C'}},
        {90, []byte{'X', 'C'}},
        {50, []byte{'L'}},
        {40, []byte{'X', 'L'}},
        {10, []byte{'X'}},
        {9, []byte{'I', 'X'}},
        {5, []byte{'V'}},
        {4, []byte{'I', 'V'}},
        {1, []byte{'I'}},
    };

    ans := []byte{}
    for _, n := range numerals {
        for num >= n.value {
            num -= n.value
            ans = append(ans, n.symbol...)
        }
    }

    return string(ans);

}
