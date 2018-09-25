#include <algorithm>
#include <iostream>
#include <map>
#include <numeric>

using namespace std;

struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
 private:
  int maxSum = numeric_limits<int>::min();
  int dfs(TreeNode* cur) {
    if (cur == nullptr) return 0;
    int left_max = max(dfs(cur->left), 0);
    int right_max = max(dfs(cur->right), 0);
    maxSum = max(maxSum, left_max + right_max + cur->val);
    return cur->val + max(left_max, right_max);
  }

 public:
  int maxPathSum(TreeNode* root) {
    dfs(root);
    return maxSum;
  }
};

int main(int argc, char const* argv[]) { return 0; }
