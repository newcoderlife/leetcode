#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<int> twoSumBasic(vector<int>& nums, int target) {
    for (int left = 0; left < nums.size(); left++) {
      for (int right = left + 1; right < nums.size(); right++) {
        if (nums[left] + nums[right] == target) {
          return vector<int>{left, right};
        }
      }
    }
    return vector<int>();
  }

  vector<int> twoSumHash(vector<int>& nums, int target) {
    unordered_map<int, int> hash;
    for (int i = 0; i < nums.size(); i++) {
      auto item = hash.find(target - nums[i]);
      if (item != hash.end()) {
        return vector<int>{item->second, i};
      }
      hash[nums[i]] = i;
    }
    return vector<int>();
  }
};