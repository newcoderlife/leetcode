impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let str0 = &strs[0].as_bytes();
        for i in 0..str0.len() {
            for s in strs.iter() {
                let str = s.as_bytes();
                if str.len() == i || str0[i] != str[i] {
                    return String::from_utf8(str0[..i].to_vec()).unwrap();
                }
            }
        }
        return strs[0].to_string();
    }
}

struct Solution {}

fn main() {
    let words = ["flower", "flow", "flight"]
        .iter()
        .map(|&s| s.into())
        .collect();
    println!(
        "common prefix is: {}",
        Solution::longest_common_prefix(words)
    );
}
