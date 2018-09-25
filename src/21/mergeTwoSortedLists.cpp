#include <iostream>

using namespace std;

struct ListNode {
  int val;
  ListNode* next;
  ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
 public:
  ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    ListNode *root = new ListNode(0), *cur_l1 = l1, *cur_l2 = l2,
             *cur_rst = root;

    while (cur_l1 && cur_l2) {
      if (cur_l1->val < cur_l2->val) {
        cur_rst->next = new ListNode(cur_l1->val);
        cur_l1 = cur_l1->next;
      } else {
        cur_rst->next = new ListNode(cur_l2->val);
        cur_l2 = cur_l2->next;
      }
      cur_rst = cur_rst->next;
    }

    for (; cur_l1; cur_l1 = cur_l1->next, cur_rst = cur_rst->next)
      cur_rst->next = new ListNode(cur_l1->val);

    for (; cur_l2; cur_l2 = cur_l2->next, cur_rst = cur_rst->next)
      cur_rst->next = new ListNode(cur_l2->val);

    return root->next;
  }
};

int main(int argc, char const* argv[]) {
  Solution sol;
  return 0;
}
