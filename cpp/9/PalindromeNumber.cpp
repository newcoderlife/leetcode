#include <string>
#include <iostream>
#include <algorithm>

using namespace std;

class Solution
{
public:
    bool isPalindrome(int x)
    {
        string num = to_string(x), reversed = num;
        reverse(reversed.begin(), reversed.end());
        return num == reversed;
    }
};

int main(int argc, char const *argv[])
{
    Solution sol;
    cout << sol.isPalindrome(131) << endl;
    return 0;
}
