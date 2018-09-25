#include <stack>
#include <vector>

using namespace std;

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
 public:
  vector<int> preorderTraversal(TreeNode *root) {
    vector<int> result;
    if (root == nullptr) return result;

    stack<TreeNode *> st;
    st.push(root);

    while (!st.empty()) {
      TreeNode *cur = st.top();
      st.pop();
      result.push_back(cur->val);

      if (cur->right) {
        st.push(cur->right);
      }
      if (cur->left) {
        st.push(cur->left);
      }
    }

    return result;
  }
};