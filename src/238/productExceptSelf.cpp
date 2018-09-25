#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<int> productExceptSelf(vector<int>& nums) {
    vector<int> results(nums.size(), 1);

    for (int i = 1; i < nums.size(); i++)
      results[i] = nums[i - 1] * results[i - 1];

    int L = 1;
    for (int i = nums.size() - 1; i >= 0; i--) {
      results[i] *= L;
      L *= nums[i];
    }

    return results;
  }
};

int main(int argc, char const* argv[]) { return 0; }
