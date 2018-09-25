#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int superEggDrop(int K, int N) {
    vector<int> dp(K + 1, 0);
    int result = 0;
    while (dp[K] < N) {
      for (int i = K; i > 0; i--) dp[i] = dp[i] + dp[i - 1] + 1;
      result++;
    }
    return result;
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().superEggDrop(1, 2) << endl;
  cout << Solution().superEggDrop(2, 6) << endl;
  cout << Solution().superEggDrop(3, 14) << endl;
  cout << Solution().superEggDrop(100, 10000) << endl;
  return 0;
}
