#include <iostream>
#include <sstream>
#include <string>
#include <vector>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
 public:
  bool hasCycle(ListNode *head) {
    if (head == nullptr) return false;

    auto fast = head->next, slow = head;
    while (fast != slow) {
      if (fast == nullptr || fast->next == nullptr) return false;
      slow = slow->next;
      fast = fast->next->next;
    }
    return true;
  }
};
