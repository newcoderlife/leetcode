#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int removeDuplicates(vector<int>& nums) {
    if (nums.size() < 2) return nums.size();
    int left = 0, right = 0;
    while (++right < nums.size())
      if (nums[left] != nums[right]) nums[++left] = nums[right];
    return left + 1;
  }
};

int main(int argc, char const* argv[]) {
  auto nums = vector<int>{1, 2};
  cout << Solution().removeDuplicates(nums) << endl;
  for (auto num : nums) {
    cout << num << ' ';
  }
  cout << endl;
  return 0;
}
