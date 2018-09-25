#include <iostream>
#include <vector>

using namespace std;

class Solution {
 private:
  vector<vector<int>> results;
  vector<bool> exists;

 public:
  vector<vector<int>> permute(vector<int>& nums) {
    exists.resize(nums.size(), false);
    vector<int> cur;
    dfs(cur, nums);
    return results;
  }

  void dfs(vector<int>& cur, vector<int>& nums) {
    if (cur.size() == nums.size()) {
      results.push_back(cur);
      return;
    }

    for (int i = 0; i < nums.size(); i++) {
      if (!exists[i]) {
        cur.push_back(nums[i]);
        exists[i] = true;
        dfs(cur, nums);
        exists[i] = false;
        cur.pop_back();
      }
    }
  }
};

int main(int argc, char const* argv[]) {
  auto results = Solution().permute(vector<int>{1, 2, 3, 4, 5, 6, 7});
  return 0;
}
