#[allow(dead_code)]
const SIMPLE_INPUT: &str = include_str!("./../../inputs/day-02/simple.txt");
#[allow(dead_code)]
const FULL_INPUT: &str = include_str!("./../../inputs/day-02/full.txt");

fn main() {
    let input = read_input(FULL_INPUT);
    println!("Result for star 1 is {}", star1(&input));
    println!("Result for star 2 is {}", star2(&input));
}

fn star1(input: &[Vec<i32>]) -> usize {
    input.iter().filter(|x| is_safe(x)).count()
}

fn star2(input: &[Vec<i32>]) -> usize {
    input
        .iter()
        .filter(|&seq| {
            let unsafe_idx = find_unsafe_idx(seq);
            match unsafe_idx {
                UnsafeResult::Safe => true,
                UnsafeResult::Unfixable => false,
                UnsafeResult::Retry(idx) => {
                    // Try changing the first
                    let mut seq_copy = seq.clone();
                    let removed_el = seq_copy.remove(idx - 1);
                    if is_safe(&seq_copy) {
                        return true;
                    }
                    // Try changing the second
                    seq_copy[idx - 1] = removed_el;
                    if is_safe(&seq_copy) {
                        return true;
                    }
                    // Try changing the first if it makes sense
                    if idx == 2 {
                        seq_copy.clone_from_slice(&seq[1..]);
                        is_safe(&seq_copy)
                    } else {
                        false
                    }
                }
            }
        })
        .count()
}

fn read_input(input: &str) -> Vec<Vec<i32>> {
    input
        .trim_end()
        .lines()
        .map(|line| {
            line.trim()
                .split_whitespace()
                .map(|x| x.parse().unwrap())
                .collect()
        })
        .collect()
}

fn is_safe(seq: &[i32]) -> bool {
    find_unsafe_idx(seq).is_safe()
}

enum UnsafeResult {
    Safe,
    Retry(usize),
    Unfixable,
}
impl UnsafeResult {
    fn is_safe(&self) -> bool {
        match self {
            UnsafeResult::Safe => true,
            _ => false,
        }
    }
}
fn find_unsafe_idx(seq: &[i32]) -> UnsafeResult {
    let first_diff = seq[1] - seq[0];
    let pos = first_diff > 0; // 0 means unsafe anyway
    for (idx, &el) in seq.iter().enumerate().skip(1) {
        let diff = el - seq[idx - 1];
        // Unfixable short path
        if idx > 1 && idx < seq.len() - 2 {
            if pos && diff > 3 {
                return UnsafeResult::Unfixable;
            } else if !pos && diff < -3 {
                return UnsafeResult::Unfixable;
            }
        }
        // Retry path
        if pos && (diff < 1 || diff > 3) {
            return UnsafeResult::Retry(idx);
        } else if !pos && (diff < -3 || diff > -1) {
            return UnsafeResult::Retry(idx);
        }
    }
    UnsafeResult::Safe
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn simple_star1() {
        let input = read_input(SIMPLE_INPUT);
        let result = star1(&input);
        assert_eq!(result, 2);
    }

    #[test]
    fn full_star1() {
        let input = read_input(FULL_INPUT);
        let result = star1(&input);
        assert_eq!(result, 680);
    }

    #[test]
    fn simple_star2() {
        let input = read_input(SIMPLE_INPUT);
        let result = star2(&input);
        assert_eq!(result, 4);
    }

    #[test]
    fn full_star2() {
        let input = read_input(FULL_INPUT);
        let result = star2(&input);
        assert_eq!(result, 710);
    }
}
