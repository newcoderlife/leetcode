#include <iostream>
#include <string>

using namespace std;

class Solution {
 public:
  string multiply(string num1, string num2) {
    if (num1.size() > num2.size()) return multiply(num2, num1);
    reverse(num1.begin(), num1.end());
    reverse(num2.begin(), num2.end());
    string result;
    result.resize(num1.size() + num2.size(), 0);

    for (int i = 0; i < num1.size(); i++) {
      for (int j = 0; j < num2.size(); j++) {
        result[i + j] += (num1[i] - '0') * (num2[j] - '0');
        if (result[i + j] > 9) {
          result[i + j + 1] += result[i + j] / 10;
          result[i + j] = result[i + j] % 10;
        }
      }
    }
    while (result.back() == 0 && result.size() != 1) result.pop_back();
    for (int i = 0; i < result.size(); i++) result[i] += '0';
    reverse(result.begin(), result.end());
    return result;
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().multiply("9", "9") << endl;
  return 0;
}
