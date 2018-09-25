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
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
    {
        ListNode *iter1 = l1, *iter2 = l2, *headPrev = new ListNode(0), *prev = headPrev;
        bool carry = false;

        while (iter1 != nullptr || iter2 != nullptr || carry)
        {
            int total = (iter1 == nullptr ? 0 : iter1->val);
            total += (iter2 == nullptr ? 0 : iter2->val);
            total += carry ? 1 : 0;

            if (total < 10)
            {
                prev->next = new ListNode(total);
                carry = false;
            }
            else
            {
                prev->next = new ListNode(total - 10);
                carry = true;
            }

            iter1 = (iter1 == nullptr ? nullptr : iter1->next);
            iter2 = (iter2 == nullptr ? nullptr : iter2->next);
            prev = prev->next;
        }

        return headPrev->next; // Memory leak here.
    }
};

ListNode *createList(const vector<int> &nums)
{
    ListNode *head = new ListNode(0), *prev = head;
    for (int i = 0; i < nums.size(); i++)
    {
        prev->val = nums[i];

        if (i + 1 != nums.size())
        {
            prev->next = new ListNode(0);
            prev = prev->next;
        }
    }

    return head;
}

int main(int argc, char *argv[])
{
    Solution sol;
    auto num1 = createList(vector<int>{1, 2, 3});
    auto num2 = createList(vector<int>{4, 5, 6});
    auto result = sol.addTwoNumbers(num1, num2);
    return 0;
}