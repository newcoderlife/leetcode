#include <iostream>
#include <vector>

using namespace std;

struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
 private:
  TreeNode* dfs(vector<int>& postorder, int& cur, vector<int>& inorder,
                int left, int right) {
    if (cur < 0) return nullptr;

    int index;
    for (index = right - 1; index >= left && postorder[cur] != inorder[index];
         index--)
      ;

    TreeNode* root = new TreeNode(postorder[cur]);
    if (index + 1 < right)
      root->right = dfs(postorder, --cur, inorder, index + 1, right);
    if (left < index) root->left = dfs(postorder, --cur, inorder, left, index);
    return root;
  }

 public:
  TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
    int cur = postorder.size() - 1;
    return dfs(postorder, cur, inorder, 0, inorder.size());
  }
};

int main(int argc, char const* argv[]) {
  auto root = Solution().buildTree(vector<int>{9, 3, 15, 20, 7},
                                   vector<int>{9, 15, 7, 20, 3});
  return 0;
}
