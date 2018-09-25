#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int maxProfit(vector<int>& prices) {
    if (prices.size() == 0) return 0;
    int result = 0, min_price = prices[0];
    for (int i = 1; i < prices.size(); i++) {
      result = max(result, prices[i] - min_price);
      min_price = min(min_price, prices[i]);
    }
    return result;
  }
};

int main(int argc, char const* argv[]) {
  vector<int> prices = {7, 1, 5, 3, 6, 4};
  cout << Solution().maxProfit(prices) << endl;
  return 0;
}
