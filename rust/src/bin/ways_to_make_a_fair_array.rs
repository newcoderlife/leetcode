impl Solution {
    pub fn ways_to_make_fair(nums: Vec<i32>) -> i32 {
        if nums.len() < 2 {
            return 1;
        }

        let mut result = 0;
        let mut even = Vec::new();
        let mut odd = Vec::new();
        for i in 0..nums.len() {
            if i & 1 == 1 {
                if odd.len() == 0 {
                    odd.push(nums[i])
                } else {
                    odd.push(nums[i] + odd[odd.len() - 1])
                }
            } else {
                if even.len() == 0 {
                    even.push(nums[i])
                } else {
                    even.push(nums[i] + even[even.len() - 1])
                }
            }
        }

        let mut pre_odd;
        let mut pre_even;
        let mut fol_odd;
        let mut fol_even;
        for i in 0..nums.len() {
            if i & 1 == 1 {
                pre_odd = odd[i / 2] - nums[i];
                pre_even = even[i / 2];
                fol_odd = even[even.len() - 1] - pre_even;
                fol_even = odd[odd.len() - 1] - odd[i / 2];
            } else {
                pre_odd = if i < 2 { 0 } else { odd[i / 2 - 1] };
                pre_even = even[i / 2] - nums[i];
                fol_odd = even[even.len() - 1] - even[i / 2];
                fol_even = odd[odd.len() - 1] - pre_odd;
            }
            if pre_even + fol_even == pre_odd + fol_odd {
                result += 1;
            }
        }

        return result;
    }
}

struct Solution {}

fn main() {
    dbg!(Solution::ways_to_make_fair(vec![2, 1, 6, 4])); // 1
    dbg!(Solution::ways_to_make_fair(vec![1, 1, 1])); // 3
    dbg!(Solution::ways_to_make_fair(vec![1, 2, 3])); // 0
}
