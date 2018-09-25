#include <iostream>
#include <stack>
#include <string>

using namespace std;

class Solution {
 public:
  bool isValid(string s) {
    stack<char> st;
    for (auto c : s) {
      if (c == '(' || c == '[' || c == '{')
        st.push(c);
      else {
        if (st.empty()) return false;
        char prev = st.top();
        st.pop();
        if ((c == ')' && prev == '(') || (c == ']' && prev == '[') ||
            (c == '}' && prev == '{'))
          continue;
        else
          return false;
      }
    }
    return st.empty();
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().isValid("()") << endl;
  return 0;
}
