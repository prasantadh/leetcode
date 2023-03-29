func mySqrt(x int) int {
   if x < 2 {
       return x
   } 

   l, h := 1, x
   for l != h - 1 {
       m := (l + h) / 2
       if m * m <= x {
           l = m
       } else {
           h = m
       }
   }
   return l
}
