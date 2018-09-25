#include <string>
#include <map>
#include <iostream>
#include <algorithm>

using namespace std;

class Solution
{
public:
    int lengthOfLongestSubstring(string s)
    {
        int index[256];
        memset(index, 0, sizeof(index));
        int result = 0;

        for (int i = 0, j = 0, len = s.length(); i < len && j < len; j++)
        {
            i = max(i, index[s[j]]);
            result = max(result, j - i + 1);
            index[s[j]] = j + 1;
        }

        return result;
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    cout << sol.lengthOfLongestSubstring("bbbbb") << endl;
    return 0;
}