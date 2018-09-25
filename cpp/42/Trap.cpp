#include <vector>
#include <algorithm>
#include <stack>
#include <iostream>

using namespace std;

class Solution
{
public:
    int trap(vector<int> &height)
    {
        int result = 0;
        stack<int> st;

        for (int i = 0; i < height.size(); i++)
        {
            while (!st.empty() && height[i] > height[st.top()])
            {
                int top = st.top();
                st.pop();
                if (st.empty())
                    break;

                int dist = i - st.top() - 1;
                int topHeight = min(height[i], height[st.top()]) - height[top];
                result += topHeight * dist;
            }
            st.push(i);
        }
        return result;
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    cout << sol.trap(vector<int>{3, 2, 1, 0, 10}) << endl;
    return 0;
}