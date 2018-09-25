#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int findDuplicate(vector<int>& nums) {
    /*
    int left = 1, right = nums.size() - 1;
    while (left < right) {
      int mid = (left + right) / 2, cnt = 0;
      for (int num : nums)
        if (num <= mid) cnt++;

      if (cnt > mid)
        right = mid;
      else
        left = mid + 1;
    }
    return left;
    */
    int fast = nums.front(), slow = nums.front();
    do {
      fast = nums[nums[fast]];
      slow = nums[slow];
    } while (fast != slow);

    fast = nums.front();
    while (fast != slow) {
      fast = nums[fast];
      slow = nums[slow];
    }

    return slow;
  }
};

int main(int argc, char const* argv[]) {
  vector<int> nums = {1, 3, 4, 2, 2};
  cout << Solution().findDuplicate(nums) << endl;
  return 0;
}
