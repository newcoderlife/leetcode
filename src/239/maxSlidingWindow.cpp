#include <deque>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
 private:
  deque<int> dq;

  void push(int num) {
    while (!dq.empty() && dq.back() < num) dq.pop_back();
    dq.push_back(num);
  }

  void pop(int num) {
    if (!dq.empty() && dq.front() == num) dq.pop_front();
  }

 public:
  vector<int> maxSlidingWindow(vector<int>& nums, int k) {
    int len = nums.size();
    vector<int> result;
    if (len * k == 0) return result;

    for (int i = 0; i < nums.size(); i++) {
      if (i < k - 1) {
        push(nums[i]);
      } else {
        push(nums[i]);
        result.push_back(dq.front());
        pop(nums[i - k + 1]);
      }
    }

    return result;
  }
};

int main(int argc, char const* argv[]) {
  auto result = Solution().maxSlidingWindow();
  return 0;
}
