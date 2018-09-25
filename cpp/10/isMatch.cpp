#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
 private:
  string getToken(string p, int cur_p) {
    if (cur_p + 1 < p.size() && p[cur_p + 1] == '*') return p.substr(cur_p, 2);
    return p.substr(cur_p, 1);
  }

  bool dfs(string s, int cur_s, string p, int cur_p) {
    if (cur_p == p.size()) return cur_s == s.size();

    string token = getToken(p, cur_p);
    if (token.size() == 2) {
      if (token.front() == '.') {
        if (cur_p + 2 == p.size()) return true;

        string next_token;
        int next_p;
        for (next_p = cur_p + 2; next_p < p.size(); next_p += 2) {
          next_token = getToken(p, next_p);
          if (next_token.size() == 1) break;
        }
        if (next_p == p.size()) return true;

        for (int i = cur_s; i < s.size(); i++)
          if ((next_token[0] == '.' || s[i] == next_token[0]) &&
              dfs(s, i + 1, p, next_p + 1))
            return true;
        return false;
      } else {
        if (dfs(s, cur_s, p, cur_p + 2)) return true;
        for (int i = cur_s; i < s.size() && s[i] == token.front(); i++)
          if (dfs(s, i + 1, p, cur_p + 2)) return true;
        return false;
      }
    } else {
      if (cur_s == s.size()) return false;
      if (token[0] == '.' || s[cur_s] == token[0])
        return dfs(s, cur_s + 1, p, cur_p + 1);
      return false;
    }
  }

  vector<vector<bool>> mem;

  bool dp(int cur_s, int cur_p, string s, string p) {
    if (mem[cur_s][cur_p] != -1) return mem[cur_s][cur_p];
    if (cur_p == p.size()) return mem[cur_s][cur_p] = (cur_s == s.size());

    bool match = cur_s < s.size() && (s[cur_s] == p[cur_p] || p[cur_p] == '.');
    if (cur_p + 1 < p.size() && p[cur_p + 1] == '*')
      return mem[cur_s][cur_p] = dp(cur_s, cur_p + 2, s, p) ||
                                 match && dp(cur_s + 1, cur_p, s, p);
    else
      return mem[cur_s][cur_p] = match && dp(cur_s + 1, cur_p + 1, s, p);
  }

 public:
  bool isMatch(string s, string p) {
    mem.resize(s.size() + 1, vector<bool>(p.size() + 1, false));
    mem[s.size()][p.size()] = 1;

    for (int i = s.size(); i >= 0; i--) {
      for (int j = p.size() - 1; j >= 0; j--) {
        bool match = (i < s.size() && (p[j] == s[i] || p[j] == '.'));
        if (j + 1 < p.size() && p[j + 1] == '*')
          mem[i][j] = mem[i][j + 2] || match && mem[i + 1][j];
        else
          mem[i][j] = match && mem[i + 1][j + 1];
      }
    }

    return mem[0][0];
    // return dp(0, 0, s, p);
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().isMatch("aa", "a") << endl;
  cout << Solution().isMatch("aa", "a*") << endl;
  cout << Solution().isMatch("ab", ".*") << endl;
  cout << Solution().isMatch("aab", "c*a*b") << endl;
  cout << Solution().isMatch("mississippi", "mis*is*p*.") << endl;
  cout << Solution().isMatch("ab", ".*c") << endl;
  cout << Solution().isMatch("bbbba", ".*a*a") << endl;
  cout << Solution().isMatch("abbbcd", "ab*bbbcd") << endl;
  cout << Solution().isMatch("", "..*") << endl;
  return 0;
}
