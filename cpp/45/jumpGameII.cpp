#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int jump(vector<int>& nums) {
    if (nums.size() < 2) return 0;
    int steps = 0;
    for (int i = 0; i < nums.size(); steps++) {
      if (i + nums[i] + 1 >= nums.size()) return steps + 1;

      int next_idx = 0, furthest = 0;
      for (int j = 1; j <= nums[i]; j++) {
        if (i + j + nums[i + j] > furthest) {
          furthest = i + j + nums[i + j];
          next_idx = i + j;
        }
      }
      i = next_idx;
    }

    return steps;
  }
};

int main(int argc, char const* argv[]) {
  vector<int> jumps{4, 1, 1, 3, 1, 1, 1};
  cout << Solution().jump(jumps) << endl;
  return 0;
}
