class Solution {
 public:
  bool isPowerOfTwo(int n) {
    /*
    while (n > 1) {
      if (n & 1 == 1) return false;
      n >>= 1;
    }
    return n == 1;
    */
    if (n == 0) return false;
    long long x = n;
    return (x & -x) == x;
  }
};
