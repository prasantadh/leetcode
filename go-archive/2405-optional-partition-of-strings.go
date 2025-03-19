func partitionString(s string) int {
  /* strategy: when a repetition is seen,
  flush the prefix, add 1 to the answer,
  start over */
   
   ans := 0
   lastseen := [128]int{}
   for i := range lastseen {
       lastseen[i] = -1
   }
   tail := 0
   lastseen[s[0]] = 0
   for i:= 1; i < len(s); i++ {
       c := s[i]

       if lastseen[c] >= tail {
           ans++
           tail = i
       }
       lastseen[c] = i
    } 
    ans++
    return ans
}
