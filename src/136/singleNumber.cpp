#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int singleNumber(vector<int>& nums) {
    int result = 0;
    for (auto num : nums) result ^= num;
    return result;
  }
};

int main(int argc, char const* argv[]) {
  auto nums = vector<int>{2, 2, 1};
  cout << Solution().singleNumber(nums) << endl;
  return 0;
}
