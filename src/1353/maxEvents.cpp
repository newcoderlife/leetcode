#include <algorithm>
#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
 public:
  int maxEvents(vector<vector<int>>& events) {
    sort(events.begin(), events.end(), [](const auto& lhs, const auto& rhs) {
      return lhs.back() < rhs.back();
    });
    unordered_set<int> exists;
    for (int i = 0; i < events.size(); i++) {
      for (int j = events[i].front(); j <= events[i].back(); j++) {
        if (exists.insert(j).second == true) break;
      }
    }
    return exists.size();
  }
};

int main(int argc, char const* argv[]) { return 0; }
