#include <iostream>
#include <vector>

using namespace std;

struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

class Solution
{
private:
    ListNode *merge(ListNode *left, ListNode *right)
    {
        if (left == nullptr)
            return right;
        else if (right == nullptr)
            return left;

        if (left->val < right->val)
        {
            left->next = merge(left->next, right);
            return left;
        }
        else
        {
            right->next = merge(right->next, left);
            return right;
        }
    }

    ListNode *mergeLists(vector<ListNode *> &lists, int left, int right)
    {
        if (left + 1 == right)
            return lists[left];

        int mid = (left + right) / 2;
        auto leftList = mergeLists(lists, left, mid);
        auto rightList = mergeLists(lists, mid, right);
        return merge(leftList, rightList);
    }

public:
    ListNode *mergeKLists(vector<ListNode *> &lists)
    {
        if (lists.size() == 0)
            return nullptr;
        return mergeLists(lists, 0, lists.size());
    }
};

int main(int argc, char const *argv[])
{
    Solution sol;
    return 0;
}
