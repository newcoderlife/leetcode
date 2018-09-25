struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {  // 这里直接原址骚操作
    auto head = new ListNode(), current = head;
    while (l1 != nullptr && l2 != nullptr) {
      current->next = new ListNode(l1->val + l2->val);
      if (current->val > 9) {
        current->next->val += 1;
        current->val -= 10;
      }
      current = current->next;
      l1 = l1->next;
      l2 = l2->next;
    }
    while (l1 != nullptr) {
      current->next = new ListNode(l1->val);
      if (current->val > 9) {
        current->next->val += 1;
        current->val -= 10;
      }
      current = current->next;
      l1 = l1->next;
    }
    while (l2 != nullptr) {
      current->next = new ListNode(l2->val);
      if (current->val > 9) {
        current->next->val += 1;
        current->val -= 10;
      }
      current = current->next;
      l2 = l2->next;
    }
    if (current->val > 9) {
      current->next = new ListNode(1);
      current->val -= 10;
    }
    return head->next;
  }
};