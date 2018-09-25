#include <string>
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

class Solution
{
public:
    string longestPalindrome(string s)
    {
        string after = "#";
        for_each(s.begin(), s.end(), [&after](char x) { after += x; after += '#'; });

        vector<int> radius;
        radius.resize(after.size());
        int pos = 0, count = 0, right = 0, start = 0;
        for (int i = 0; i < after.size(); i++)
        {
            radius[i] = (i < right) ? min(radius[2 * pos - i], right - i) : 1;

            while (i - radius[i] >= 0 && i + radius[i] < after.size() && after[i - radius[i]] == after[i + radius[i]])
                radius[i]++;

            if (radius[i] + i - 1 > right)
            {
                right = radius[i] + i - 1;
                pos = i;
            }

            if (count < radius[i])
            {
                count = radius[i];
                start = i - radius[i] + 1;
            }
        }

        string result;
        for_each(after.begin() + start, after.begin() + start + count * 2 - 1, [&result](char x) { if (x != '#') result+=x; });
        return result;
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    cout << sol.longestPalindrome("") << endl;
    return 0;
}