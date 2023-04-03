func minNumber(nums1 []int, nums2 []int) int {
  /* approach:
  if there is a single digit that satisfies the answer,
  use it. else create a two digit response.
  */
    memo := [10]int{}
    m1 := 10
    m2 := 10
    for _, num := range nums1 {
        memo[num] = 1
        if num < m1 {
            m1 = num
        }
    }
    for _, num := range nums2 {
        memo[num]++
        if num < m2 {
            m2 = num
        }
    }
    for i, num := range memo {
        if num == 2 {
            return i
        }
    }

    if m1 < m2 {
        return m1 * 10 + m2
    }
    return m2 * 10 + m1
}
