#include <iostream>

using namespace std;

class Solution {
 public:
  int mySqrt(int x) {
    /*
    if (x < 2) return x;

    int left = 2, right = x / 2;
    while (left <= right) {
      long long mid = left + (right - left) / 2;
      long long num = mid * mid;
      if (x < num)
        right = mid - 1;
      else if (num < x)
        left = mid + 1;
      else
        return mid;
    }
    return right;
    */
    if (x < 2) return x;
    long long left = mySqrt(x >> 2) << 1;
    long long right = left + 1;
    return (right * right > x) ? left : right;
  }
};

int main(int argc, char const *argv[]) {
  for (int i = 0; i < 100; i++) cout << i << ":" << Solution().mySqrt(i) << " ";
  cout << endl;
  return 0;
}
