#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
 private:
  vector<string> results;

  bool judge(const string &s) {
    if (s.size() > 3) return false;
    int num = stoi(s);
    if (num > 255) return false;
    if (s.size() > 1 && s[0] == '0') return false;
    return true;
  }

  void process(int i, int j, int k, const string s) {
    string a, b, c, d;
    a = s.substr(0, i);
    b = s.substr(i, j - i);
    c = s.substr(j, k - j);
    d = s.substr(k, s.size() - k);
    if (judge(a) && judge(b) && judge(c) && judge(d))
      results.push_back(a + '.' + b + '.' + c + '.' + d);
  }

 public:
  vector<string> restoreIpAddresses(string s) {
    if (s.size() > 12 || s.size() < 4) return results;
    for (int i = 1; i < 4; i++)
      for (int j = i + 1; j < min<int>(i + 4, s.size()); j++)
        for (int k = j + 1; k < min<int>(j + 4, s.size()); k++)
          process(i, j, k, s);
    return results;
  }
};

int main(int argc, char const *argv[]) {
  Solution sol;
  auto result = sol.restoreIpAddresses("010010");
  return 0;
}
