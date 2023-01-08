impl Solution {
    pub fn prefix_count(words: Vec<String>, pref: String) -> i32 {
        return words.iter().filter(|w| w.starts_with(&pref)).count() as i32;
    }
}

struct Solution {}

fn main() {
    let words = ["pay", "attention", "practice", "attend"]
        .iter()
        .map(|&s| s.into())
        .collect();
    let perf = "at".to_string();
    println!("{}", Solution::prefix_count(words, perf));
}
