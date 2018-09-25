#include <iostream>
#include <vector>

using namespace std;

struct ListNode {
  int val;
  ListNode* next;
  ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
 public:
  ListNode* removeNthFromEnd(ListNode* head, int n) {
    ListNode* prev = new ListNode(0);
    prev->next = head;
    ListNode *cur = prev, *node = prev;
    for (int i = 0; i < n; i++) cur = cur->next;

    while (cur->next != nullptr) {
      prev = prev->next;
      cur = cur->next;
    }

    ListNode* tmp = prev->next->next;
    delete prev->next;
    prev->next = tmp;

    return node->next;
  }
};

int main(int argc, char const* argv[]) { return 0; }
