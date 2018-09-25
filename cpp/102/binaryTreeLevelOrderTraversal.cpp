#include <queue>
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
  vector<vector<int>> levelOrder(TreeNode *root) {
    vector<vector<int>> results;
    if (root == nullptr) return results;

    queue<pair<TreeNode *, int>> Q;
    Q.push(make_pair(root, 0));

    while (!Q.empty()) {
      auto [now, depth] = Q.front();
      Q.pop();

      if (results.size() == depth)
        results.push_back({now->val});
      else
        results[depth].push_back(now->val);

      if (now->left) Q.push(make_pair(now->left, depth + 1));
      if (now->right) Q.push(make_pair(now->right, depth + 1));
    }

    return results;
  }
};

int main(int argc, char const *argv[]) { return 0; }
