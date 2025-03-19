#include "tags.h"
#include <cassert>
#include <iostream>
#include <unordered_map>

using namespace std;

vector<Tags> tags = {Tags::ARRAYS_AND_HASHING};

// reason about the correctness of the algorithm: loop invariant
// runtime complexity: O(N) runs
// better complexity possible? No
// memory complexity: O(V) where V is the maximum value in the array
// alternative: O(nlogn) possible with sorting
class Solution {
public:
  bool containsDuplicate(vector<int> &nums) {
    unordered_map<int, bool> attendance;
    // loop invariant: attendance[num] is true if num is in nums
    for (auto num : nums) {
      if (attendance[num])
        return true;
      attendance[num] = true;
    }
    return false;
  }
};

int main() {
  cout << "solution to 217 contains duplicate" << endl;
  vector<int> input1 = {1, 2, 3, 4, 5};
  vector<int> input2 = {1, 2, 3, 4, 5, 1};
  vector<int> input3 = {-1, -1, 2, 3, 4, 5, 1};

  Solution solution = Solution();
  assert(solution.containsDuplicate(input1) == false);
  assert(solution.containsDuplicate(input2) == true);
  assert(solution.containsDuplicate(input3) == true);
  return 0;
}
