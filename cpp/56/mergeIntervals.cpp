#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> merge(vector<vector<int>>& intervals) {
    vector<vector<int>> results;
    if (intervals.size() == 0) return results;

    sort(intervals.begin(), intervals.end(),
         [](const auto& left, const auto& right) -> bool {
           if (left.front() == right.front())
             return left.back() < right.back();
           else
             return left.front() < right.front();
         });

    for (int i = 0, j = 0; i < intervals.size(); i = j) {
      int left = intervals[i].front(), right = intervals[i].back();
      for (j = i + 1; j < intervals.size(); j++) {
        if (intervals[j].front() > right) break;
        right = max(intervals[j].back(), right);
      }
      results.push_back({left, right});
    }
    return results;
  }
};

int main(int argc, char const* argv[]) {
  Solution sol;
  vector<vector<int>> intervals{{1, 3}, {2, 6}, {8, 10}, {15, 18}};
  auto results = sol.merge(intervals);
  for (const auto& result : results)
    cout << result.front() << ", " << result.back() << endl;
  return 0;
}
