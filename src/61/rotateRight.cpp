#include <algorithm>
#include <iostream>
#include <sstream>
#include <vector>

using namespace std;

struct ListNode {
  int val;
  ListNode* next;
  ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
 private:
  auto iterList(ListNode* head) {
    if (head == nullptr) return pair(head, 0);

    int result = 1;
    ListNode* cur = head;
    while (cur->next) {
      cur = cur->next;
      result++;
    }
    cur->next = head;

    return pair(cur, result);
  }

  auto handleResultHead(ListNode* head, int k, int len) {
    ListNode *result = head, *prev = nullptr;
    for (int i = k; i < len; i++, prev = result, result = result->next)
      ;

    return pair(prev, result);
  }

 public:
  ListNode* rotateRight(ListNode* head, int k) {
    if (k == 0 || head == nullptr) return head;

    auto [tail, len] = iterList(head);
    auto [prev, result] = handleResultHead(head, k % len, len);

    prev->next = nullptr;
    return result;
  }
};

void trimLeftTrailingSpaces(string& input) {
  input.erase(input.begin(), find_if(input.begin(), input.end(),
                                     [](int ch) { return !isspace(ch); }));
}

void trimRightTrailingSpaces(string& input) {
  input.erase(
      find_if(input.rbegin(), input.rend(), [](int ch) { return !isspace(ch); })
          .base(),
      input.end());
}

vector<int> stringToIntegerVector(string input) {
  vector<int> output;
  trimLeftTrailingSpaces(input);
  trimRightTrailingSpaces(input);
  input = input.substr(1, input.length() - 2);
  stringstream ss;
  ss.str(input);
  string item;
  char delim = ',';
  while (getline(ss, item, delim)) {
    output.push_back(stoi(item));
  }
  return output;
}

ListNode* stringToListNode(string input) {
  // Generate list from the input
  vector<int> list = stringToIntegerVector(input);

  // Now convert that list into linked list
  ListNode* dummyRoot = new ListNode(0);
  ListNode* ptr = dummyRoot;
  for (int item : list) {
    ptr->next = new ListNode(item);
    ptr = ptr->next;
  }
  ptr = dummyRoot->next;
  delete dummyRoot;
  return ptr;
}

int stringToInteger(string input) { return stoi(input); }

string listNodeToString(ListNode* node) {
  if (node == nullptr) {
    return "[]";
  }

  string result;
  while (node) {
    result += to_string(node->val) + ", ";
    node = node->next;
  }
  return "[" + result.substr(0, result.length() - 2) + "]";
}

int main() {
  string line;
  while (getline(cin, line)) {
    ListNode* head = stringToListNode(line);
    getline(cin, line);
    int k = stringToInteger(line);

    ListNode* ret = Solution().rotateRight(head, k);

    string out = listNodeToString(ret);
    cout << out << endl;
  }
  return 0;
}