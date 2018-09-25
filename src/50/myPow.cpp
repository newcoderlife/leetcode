#include <algorithm>
#include <cmath>
#include <iostream>

using namespace std;

class Solution {
 public:
  double myPow(double x, int n) { return pow(x, n); }
};

int main(int argc, char const *argv[]) {
  cout << Solution().myPow(2.0, -2) << endl;
  cout << Solution().myPow(-10.0, (1 << 30)) << endl;
  return 0;
}
