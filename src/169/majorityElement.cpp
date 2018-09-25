#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  int majorityElement(vector<int>& nums) {
    if (nums.empty()) return 0;
    if (nums.size() < 3) return nums[0];

    int mode = nums[0], count = 1;
    for (int i = 1; i < nums.size(); i++) {
      if (count == 0) mode = nums[i];
      count += (nums[i] == mode) ? 1 : -1;
    }
    return mode;
  }
};

int main(int argc, char const* argv[]) {
  cout << Solution().majorityElement(vector<int>{1, 2, 1, 2, 1, 1}) << endl;
  return 0;
}
