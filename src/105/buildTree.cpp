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
  TreeNode* dfs(vector<int>& preorder, int& cur, vector<int>& inorder, int left,
                int right) {
    if (cur == preorder.size()) return nullptr;

    int index;
    for (index = left; index <= right && preorder[cur] != inorder[index];
         index++)
      ;

    TreeNode* root = new TreeNode(preorder[cur]);
    if (left < index) root->left = dfs(preorder, ++cur, inorder, left, index);
    if (index + 1 < right)
      root->right = dfs(preorder, ++cur, inorder, index + 1, right);
    return root;
  }

 public:
  TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
    int cur = 0;
    return dfs(preorder, cur, inorder, 0, inorder.size());
  }
};

int main(int argc, char const* argv[]) {
  auto root = Solution().buildTree(vector<int>{3, 9, 20, 15, 7},
                                   vector<int>{9, 3, 15, 20, 7});
  return 0;
}
