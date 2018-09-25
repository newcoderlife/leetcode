#include <iostream>
#include <vector>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(NULL) {}
};

ListNode *vec2List(vector<int> &content, int index) {
  if (index == content.size()) return nullptr;

  ListNode *head = new ListNode(content[index]);
  head->next = vec2List(content, index + 1);
  return head;
}

class Solution {
 public:
  ListNode *reverseList(ListNode *head) {
    ListNode *prev = nullptr;
    for (ListNode *tmp; head; head = tmp) {
      tmp = head->next;
      head->next = prev;
      prev = head;
    }
    return prev;
  }
};

int main(int argc, char *argv[]) {
  Solution sol;
  vector<int> tmp = {1, 2, 3};
  auto result = sol.reverseList(vec2List(tmp, 0));
  for (ListNode *head = result; head != nullptr; head = head->next)
    cout << head->val << endl;
  return 0;
}