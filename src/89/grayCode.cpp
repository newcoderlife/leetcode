#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<int> grayCode(int n) {
    vector<int> results = {0};

    for (int i = 0, flag = 1, len; i < n; i++, flag <<= 1)
      for (int j = results.size() - 1; j >= 0; j--)
        results.push_back(results[j] | flag);

    return results;
  }
};

int main(int argc, char const *argv[]) {
  auto results = Solution().grayCode(5);
  for (auto i : results) cout << i << endl;
  return 0;
}
