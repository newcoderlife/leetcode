#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int search(vector<int>& nums, int target) {
    int left = 0, right = nums.size() - 1;
    while (left < right) {
      int pivot = (left + right) / 2;
      if ((nums[0] > target) ^ (nums[0] > nums[pivot]) ^ (target > nums[pivot]))
        left = pivot + 1;
      else
        right = pivot;
    }

    return left == right && nums[left] == target ? left : -1;
  }
};

int main(int argc, char const* argv[]) {
  Solution sol;
  vector<int> nums{1, 3};

  cout << sol.search(nums, 1) << endl;
  return 0;
}
