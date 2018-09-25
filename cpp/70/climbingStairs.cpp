#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int climbStairs(int n) {
    if (n < 2) return 1;
    int prev = 1, now = 1;
    for (int i = 1; i < n; i++) {
      int tmp = prev + now;
      prev = now;
      now = tmp;
    }
    return now;
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().climbStairs(3) << endl;
  return 0;
}
