#include <algorithm>
#include <cmath>
#include <iostream>
#include <numeric>
#include <sstream>
#include <vector>

using namespace std;

class Solution {
 public:
  int threeSumClosest(vector<int> &nums, int target) {
    if (nums.size() < 4) return accumulate(nums.begin(), nums.end(), 0);

    sort(nums.begin(), nums.end());

    int result = nums[0] + nums[1] + nums[2];
    for (int i = 0; i + 2 < nums.size(); i++) {
      int left = i + 1, right = nums.size() - 1;
      while (left < right) {
        int sum = nums[i] + nums[left] + nums[right];
        if (abs(sum - target) < abs(result - target)) result = sum;

        if (sum > target)
          right--;
        else if (sum < target)
          left++;
        else
          return sum;
      }
    }

    return result;
  }
};

void trimLeftTrailingSpaces(string &input) {
  input.erase(input.begin(), find_if(input.begin(), input.end(),
                                     [](int ch) { return !isspace(ch); }));
}

void trimRightTrailingSpaces(string &input) {
  input.erase(
      find_if(input.rbegin(), input.rend(), [](int ch) { return !isspace(ch); })
          .base(),
      input.end());
}

vector<int> stringToIntegerVector(string input) {
  vector<int> output;
  trimLeftTrailingSpaces(input);
  trimRightTrailingSpaces(input);
  input = input.substr(1, input.length() - 2);
  stringstream ss;
  ss.str(input);
  string item;
  char delim = ',';
  while (getline(ss, item, delim)) {
    output.push_back(stoi(item));
  }
  return output;
}

int stringToInteger(string input) { return stoi(input); }

int main() {
  string line;
  while (getline(cin, line)) {
    vector<int> nums = stringToIntegerVector(line);
    getline(cin, line);
    int target = stringToInteger(line);

    int ret = Solution().threeSumClosest(nums, target);

    string out = to_string(ret);
    cout << out << endl;
  }
  return 0;
}