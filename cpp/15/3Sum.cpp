#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> threeSum(vector<int>& nums) {
    vector<vector<int>> results;
    if (nums.empty()) return results;
    sort(nums.begin(), nums.end());
    if (nums.front() > 0 || nums.back() < 0) return results;

    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] > 0) break;
      if (i > 0 && nums[i] == nums[i - 1]) continue;

      int left = i + 1, right = nums.size() - 1;
      while (left < right) {
        int sum = nums[i] + nums[left] + nums[right];
        if (sum == 0) {
          results.push_back({nums[i], nums[left], nums[right]});
          while (left < right && nums[left] == nums[left + 1]) left++;
          while (left < right && nums[right - 1] == nums[right]) right--;
          left++, right--;
        } else if (sum < 0)
          left++;
        else
          right--;
      }
    }

    return results;
  }
};

int main(int argc, char const* argv[]) {
  Solution sol;
  vector<int> nums{0, 0, 0, 0};
  auto results = sol.threeSum(nums);
  for (const auto& result : results) {
    for (int i = 0; i < result.size(); i++) cout << result[i] << ' ';
    cout << endl;
  }
  return 0;
}
