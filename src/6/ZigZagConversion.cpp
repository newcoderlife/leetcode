#include <string>
#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
    string convert(string s, int numRows)
    {
        if (numRows == 1)
            return s;

        string result;
        for (int i = 0, step = 2 * numRows - 2; i < numRows; i++)
        {
            for (int j = 0; j + i < s.size(); j += step)
            {
                result.push_back(s[i + j]);
                if (i != 0 && i != numRows - 1 && j + step - i < s.size())
                    result.push_back(s[j + step - i]);
            }
        }

        return result;
    }
};

int main(int argc, char const *argv[])
{
    Solution sol;
    cout << sol.convert("PAYPALISHIsasdasdasdasddfaszdfgsdfgsdfgxdfgxdfgRING", 4) << endl;
    return 0;
}
