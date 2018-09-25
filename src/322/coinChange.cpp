#include <algorithm>
#include <iostream>
#include <numeric>
#include <vector>

using namespace std;

class Solution {
 public:
  int coinChange(vector<int>& coins, int amount) {
    if (amount == 0) return amount;

    sort(coins.begin(), coins.end());
    vector<int> dp(amount + 1, numeric_limits<int>::max() / 2);
    dp[0] = 0;
    for (int j = 0; j < coins.size(); j++)
      if (coins[j] <= amount) dp[coins[j]] = 1;

    for (int i = 1; i <= amount; i++)
      for (int j = 0; j < coins.size() && coins[j] <= i; j++) {
        dp[i] = min(dp[i - coins[j]] + 1, dp[i]);
      }

    return (dp[amount] >= numeric_limits<int>::max() / 2 ? -1 : dp[amount]);
  }
};

int main(int argc, char const* argv[]) {
  cout << Solution().coinChange(vector<int>{1, 2, 5}, 11) << endl;
  cout << Solution().coinChange(vector<int>{2}, 3) << endl;
  cout << Solution().coinChange(vector<int>{1}, 0) << endl;
  return 0;
}
