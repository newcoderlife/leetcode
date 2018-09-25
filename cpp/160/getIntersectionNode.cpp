#include <unordered_set>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
 public:
  /*
   ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
     unordered_set<intptr_t> exists;
     ListNode *cur1 = headA, *cur2 = headB;
     while (cur1 || cur2) {
       if (cur1) {
         if (!exists.insert((intptr_t)cur1).second) return cur1;
         cur1 = cur1->next;
       }
       if (cur2) {
         if (!exists.insert((intptr_t)cur2).second) return cur2;
         cur2 = cur2->next;
       }
     }
     return nullptr;
   }
   */

  ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
    ListNode *cur1 = headA, *cur2 = headB;
    while (cur1 != cur2) {
      if (cur1)
        cur1 = cur1->next;
      else
        cur1 = headB;

      if (cur2)
        cur2 = cur2->next;
      else
        cur2 = headA;
    }
    return cur1;
  }
};

int main(int argc, char const *argv[]) { return 0; }
