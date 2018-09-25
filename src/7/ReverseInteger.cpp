#include <iostream>
#include <algorithm>
#include <string>

using namespace std;

class Solution
{
  public:
    int reverse(int num)
    {
        long long x = num;
        bool negative = x < 0;
        x = abs(x);

        string str = to_string(x);
        std::reverse(str.begin(), str.end());

        long long result = atoll(str.c_str());
        result = negative ? -result : result;

        if (result < numeric_limits<int>::min() || result > numeric_limits<int>::max())
            return 0;
        return result;
    }
};

int main(int argc, char const *argv[])
{
    Solution sol;
    cout << sol.reverse(-100) << endl;
    return 0;
}
