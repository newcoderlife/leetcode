#include <iostream>
#include <numeric>
#include <vector>

using namespace std;

class Solution {
 public:
  int candy(vector<int>& ratings) {
    vector<int> caddies(ratings.size(), 1);

    for (int i = 1; i < ratings.size(); i++)
      if (ratings[i - 1] < ratings[i])
        caddies[i] = max(caddies[i], caddies[i - 1] + 1);

    for (int i = ratings.size() - 1; i > 0; i--)
      if (ratings[i - 1] > ratings[i])
        caddies[i - 1] = max(caddies[i - 1], caddies[i] + 1);

    int result = 0;
    for (auto caddy : caddies) result += caddy;
    return result;
  }
};

int main(int argc, char const* argv[]) { return 0; }
