#include <iostream>
#include <vector>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(nullptr) {}
};

ListNode *vec2List(vector<int> &content, int index) {
  if (index == content.size()) return nullptr;

  ListNode *head = new ListNode(content[index]);
  head->next = vec2List(content, index + 1);
  return head;
}

class Solution {
 private:
  ListNode *getMid(ListNode *head) {
    ListNode *slow = head, *fast = head, *prev = head;
    while (fast && fast->next) {
      prev = slow;
      slow = slow->next;
      fast = fast->next->next;
    }
    return prev;
  }

  ListNode *merge(ListNode *left, ListNode *right) {
    if (left == nullptr)
      return right;
    else if (right == nullptr)
      return left;

    if (left->val <= right->val) {
      left->next = merge(left->next, right);
      return left;
    } else {
      right->next = merge(right->next, left);
      return right;
    }
  }

 public:
  ListNode *sortList(ListNode *head) {
    if (!head || !head->next) return head;

    ListNode *mid = getMid(head);
    ListNode *right = sortList(mid->next);
    mid->next = nullptr;
    ListNode *left = sortList(head);

    return merge(left, right);
  }
};

int main(int argc, char const *argv[]) {
  Solution sol;
  auto head = sol.sortList(vec2List(vector<int>{4, 2, 1, 3}, 0));
  for (ListNode *now = head; now; now = now->next) cout << now->val << endl;
  return 0;
}
