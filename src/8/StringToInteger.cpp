#include <iostream>
#include <string>
#include <numeric>

using namespace std;

class Solution
{
public:
    int myAtoi(string str)
    {
        int i = 0;
        for (i = 0; i < str.size() && isspace(str[i]); i++)
            ;

        bool negative = false;
        if (str[i] == '+' || str[i] == '-')
            negative = str[i++] == '-';

        long long result = 0;
        while ('0' <= str[i] && str[i] <= '9')
        {
            result *= 10;
            result += str[i++] - '0';
            if (result > numeric_limits<int>::max())
            {
                if (negative)
                    return numeric_limits<int>::min();
                else
                    return numeric_limits<int>::max();
            }
        }

        result = negative ? -result : result;
        return result;
    }
};

int main(int argc, char const *argv[])
{
    Solution sol;
    cout << sol.myAtoi("20000000000000000000") << endl;
    return 0;
}
