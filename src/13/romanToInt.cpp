#include <algorithm>
#include <iostream>
#include <string>

using namespace std;

class Solution {
 public:
  int romanToInt(string s) {
    int result = 0;
    for (int i = 0, current; i < s.length(); i++, result += current) {
      current = getValue(s, i);
      if ((s[i] == 'I' || s[i] == 'X' || s[i] == 'C') &&
          (current < getValue(s, i + 1))) {
        current = -current;
      }
    }
    return result;
  }

 private:
  int getValue(string_view s, int pos) {
    if (pos >= s.length()) {
      return 0;
    }
    switch (s[pos]) {
      case 'I':
        return 1;
      case 'V':
        return 5;
      case 'X':
        return 10;
      case 'L':
        return 50;
      case 'C':
        return 100;
      case 'D':
        return 500;
      case 'M':
        return 1000;
      default:
        return 0;
    }
  }
};

int main(int argc, char *argv[]) {
  Solution sol;
  cout << sol.romanToInt("MCMXCIV") << endl;
  return 0;
}