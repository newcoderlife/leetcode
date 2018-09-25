#include <iostream>

using namespace std;

struct ListNode {
  int val;
  ListNode* next;
  ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
 public:
  ListNode* detectCycle(ListNode* head) {
    ListNode *fast = head, *slow = head;
    do {
      if (fast == nullptr || fast->next == nullptr) return nullptr;
      slow = slow->next;
      fast = fast->next->next;
    } while (fast != slow);

    fast = head;
    while (fast != slow) {
      fast = fast->next;
      slow = slow->next;
    }
    return fast;
  }
};
