# leetcode

practicing go. solving leetcode. 

this readme will contain primarily try-harder kinda try-X tips for go.
my background is primarily in c/c++/python and the "tips" here are likely
to reflect that

- a lambda definition is only available after declaration. for recursive lambda use:
```go
var f func (x int) int 
f = func (x int) int {
  if x < 1 {
    return 0
  } 
  return f(x - 1)
}
```
- go doesn't have isAlphaNum() sorta function for characters even though unicode.isLetter is available
- `s=strings.ToLower(s)` will lowercase a string
- sorting is available via `sort.Ints(nums)`. no return value. sorting is done in place.
