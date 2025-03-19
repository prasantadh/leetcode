const INF = 10000000

func maxSatisfaction(satisfaction []int) int {

    sort.Ints(satisfaction)
    satisfaction = append([]int{0}, satisfaction...)

    n := len(satisfaction)
    memo := [505][505]int{}

    finalans := 0
    for dish := 1; dish < n; dish++ {
        for position := 1; position <= dish; position++ {

            ans := -INF
            for i := position - 1; i < dish; i++ {
                if t := memo[i][position - 1] + satisfaction[dish] * position; t > ans {
                    ans = t
                }
            }
            memo[dish][position] = ans

            if ans > finalans {
                finalans = ans
            }

        }
    }

    return finalans

}
