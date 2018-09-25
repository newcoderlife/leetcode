# Median of Two Sorted Arrays

## Intro

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

**Example 1:**

```
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```

**Example 2:**

```
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

## Solution

我的第一想法是：中位数是建立在有序的情况下，所以对两个数组合并后排序，取得中位数。复杂度为`O(nlogn)`。

效率更高的一种做法是：两数组原本就是有序的，可以直接算出中位数在合并后数组的位置，然后逐个比较两数组（即插入排序）。这样不需要完成排序即可得到中位数。复杂度为`O(n)`。

最优解是：把 j 转化为 i，然后对 i 二分查找。
