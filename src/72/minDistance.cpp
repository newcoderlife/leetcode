#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  int minDistance(string word1, string word2) {
    if (word1.empty()) return word2.size();
    if (word2.empty()) return word1.size();

    int n = word1.size(), m = word2.size();

    vector<vector<int>> dp;
    dp.resize(n + 1, vector<int>(m + 1, 0));

    for (int i = 1; i <= n; i++) dp[i][0] = i;
    for (int i = 1; i <= m; i++) dp[0][i] = i;

    for (int i = 1; i <= n; i++)
      for (int j = 1; j <= m; j++) {
        dp[i][j] = dp[i - 1][j - 1] + (word1[i - 1] != word2[j - 1]);
        dp[i][j] = min(dp[i][j], dp[i - 1][j] + 1);
        dp[i][j] = min(dp[i][j], dp[i][j - 1] + 1);
      }

    return dp.back().back();
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().minDistance("horse", "ros") << endl;
  cout << Solution().minDistance("intention", "execution") << endl;
  cout << Solution().minDistance(
              "pneumonoultramicroscopicsilicovolcanoconiosis",
              "ultramicroscopically")
       << endl;
  cout << Solution().minDistance("", "") << endl;
  return 0;
}
