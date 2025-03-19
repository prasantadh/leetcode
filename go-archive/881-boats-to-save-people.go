func numRescueBoats(people []int, limit int) int {
  /* general approach is to 
  fit the heaviest person with the lightest person.
  if this doesn't work, noone else can be on the same boat
  as the heaviest person. give this person a boat of their own.
  continue
  */
    sort.Ints(people)
    i := 0
    j := len(people) - 1
    ans := 0
    for i <= j {
        if people[j] + people[i] <= limit {
            i++
        }
        j--
        ans++
    }
    return ans
}
