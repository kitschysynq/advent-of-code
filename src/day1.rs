use std::collections::HashSet;
use aoc_runner_derive::{aoc, aoc_generator};

#[aoc_generator(day1)]
pub fn part1_generator(input: &str) -> Vec<u64> {
    input.lines()
        .map(|i| i.parse().unwrap())
        .collect()
}

#[aoc(day1, part1)]
pub fn part1(input: &Vec<u64>) -> u64 {
    let target = 2020;
    let mut seen = HashSet::new();
    for x in input.iter() {
        let needs = target - x;
        if seen.contains(&needs) {
            return x.saturating_mul(needs);
        }
        seen.insert(x);
    }
    0
}

#[aoc(day1, part2, dumb)]
pub fn part2_dumb(input: &Vec<u64>) -> u64 {
    let target = 2020;
    for x in input.iter() {
        for y in input.iter() {
            for z in input.iter() {
                if x + y + z == target {
                    return x * y * z;
                }
            }
        }
    }
    0
}

#[aoc(day1, part2, not_as_dumb)]
pub fn part2_not_as_dumb(input: &Vec<u64>) -> u64 {
    let target = 2020;
    for x in input.iter() {
        let mut seen = HashSet::new();
        let inner_target = target - x;
        for y in input.iter() {
            let needs = inner_target - y;
            if seen.contains(&needs) {
                return x * y * needs;
            }
            seen.insert(y);
        }
    }
    0
}
