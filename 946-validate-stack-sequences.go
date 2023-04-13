func validateStackSequences(pushed []int, popped []int) bool {
    /* 
    stack is a LIFO structure
    means if a value x is popped,
    everything that was pushed until then and hasn't been popped yet,
    has to appear at a later index than x in the popped list
    */
    n := len(pushed)
    t := []int{}

    /*
    for each iteration of the loop,
    idx points to the next value to be pushed,
    and t contains the values that are pushed but not popped 
    before popping popval
    */
    idx := 0
    for _, popval := range popped {
        if l := len(t); l != 0 && t[l-1] == popval {
            t = t[:l-1]
        } else {
            for idx < n {
                if pushed[idx] == popval {
                    idx++
                    break
                }
                t = append(t, pushed[idx])
                idx++
            }
        }
    }
    return len(t) == 0
}
