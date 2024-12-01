use std::collections::HashMap;

const SIMPLE_INPUT: &str = include_str!("./../../inputs/day-01/simple.txt");
const FULL_INPUT: &str = include_str!("./../../inputs/day-01/full.txt");

fn star1(input: &(Vec<i32>, Vec<i32>)) -> usize {
    let mut left_list = input.0.clone();
    let mut right_list = input.1.clone();
    left_list.sort();
    right_list.sort();
    left_list
        .iter()
        .zip(right_list.iter())
        .map(|(left, right)| (left - right).abs() as usize)
        .sum()
}

fn star2(input: &(Vec<i32>, Vec<i32>)) -> usize {
    let left_map = count_map(&input.0);
    let right_map = count_map(&input.1);
    left_map
        .iter()
        .map(|(n, left_count)| {
            let occurrence = left_count * right_map.get(n).unwrap_or(&0);
            occurrence * (*n as usize)
        })
        .sum()
}

fn read_input(input: &str) -> (Vec<i32>, Vec<i32>) {
    input
        .trim_end()
        .lines()
        .map(|line| {
            let mut iter = line.trim().split("   ");
            (
                iter.next().unwrap().parse::<i32>().unwrap(),
                iter.next().unwrap().parse::<i32>().unwrap(),
            )
        })
        .unzip()
}

fn count_map(list: &[i32]) -> HashMap<i32, usize> {
    let mut map = HashMap::new();
    for el in list.iter() {
        let val = map.entry(*el).or_default();
        *val += 1;
    }
    map
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn simple_star1() {
        let input = read_input(SIMPLE_INPUT);
        let result = star1(&input);
        assert_eq!(result, 11);
    }

    #[test]
    fn full_star1() {
        let input = read_input(FULL_INPUT);
        let result = star1(&input);
        assert_eq!(result, 1388114);
    }

    #[test]
    fn simple_star2() {
        let input = read_input(SIMPLE_INPUT);
        let result = star2(&input);
        assert_eq!(result, 31);
    }

    #[test]
    fn full_star2() {
        let input = read_input(FULL_INPUT);
        let result = star2(&input);
        assert_eq!(result, 23529853);
    }
}
