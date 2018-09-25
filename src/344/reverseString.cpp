#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  void reverseString(vector<char> &s) {
    for (int i = 0; i * 2 < s.size(); i++) {
      iter_swap(s.begin() + i, s.end() - i - 1);
    }
  }
};

int main(int argc, char const *argv[]) {
  auto str = vector<char>{'h', 'e', 'l', 'l', 'o'};
  Solution().reverseString(str);
  for (auto ch : str) cout << ch;
  return 0;
}
