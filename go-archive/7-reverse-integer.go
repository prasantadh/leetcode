func reverse(x int) int {
  
  /* this one took somwhat surprisingly long
  - the negative integers modulus was returning negative numbers
  - the overflow had to be handled digit by digit
  */

    maxint := [10]byte{2,1,4,7,4,8,3,6,4,7}
    x_arr := []byte{}
    is_neg := x < 0
    if is_neg {
        x = -x
    }

    for x != 0 {
        x_arr = append(x_arr, byte(x % 10))
        x = x / 10
    }

    ans := 0
    if len(x_arr) == 10 {
        for i := 0; i < 10; i++ {
            if x_arr[i] < maxint[i] {
                break
            } else if x_arr[i] > maxint[i] {
                return 0
            }    
        }
    }

    for _, r := range x_arr {
        ans = ans * 10 + int(r)
    }

    if is_neg {
        return -ans
    }
    return ans

}
