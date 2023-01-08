impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut result = 0;
        for i in 0..nums.len() {
            if nums[i] == val {
                continue;
            }
            if i != result {
                nums[result] = nums[i];
            }
            result += 1;
        }
        return result as i32;
    }
}

struct Solution {}

fn main() {
    let mut nums = vec![3, 2, 2, 3];
    let val = 3;
    let result = Solution::remove_element(&mut nums, val) as usize;
    dbg!(&nums[..result]);
}
