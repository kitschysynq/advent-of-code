use aoc_runner_derive::{aoc, aoc_generator};
use regex::Regex;

pub struct Entry {
    min: usize,
    max: usize,
    letter: char,
    password: String,
}

#[aoc_generator(day2)]
pub fn generator(input: &str) -> Vec<Entry> {
    let re = Regex::new(r"(\d*)-(\d*) (\w): (\w*)").unwrap();

    input.lines()
        .map(|e| {
            let caps = re.captures(e).unwrap();
            Entry {
                min: caps.get(1).unwrap().as_str().parse().unwrap(),
                max: caps.get(2).unwrap().as_str().parse().unwrap(),
                letter: caps.get(3).unwrap().as_str().chars().next().unwrap(),
                password: caps.get(4).unwrap().as_str().into(),
            }
        })
        .collect()
}

#[aoc(day2, part1)]
pub fn part1(input: &Vec<Entry>) -> usize {
    let mut count = 0;
    for e in input.iter() {
        let num = e.password.chars().filter(|c| *c == e.letter).count();
        if num >= e.min && num <= e.max {
            count += 1;
        }
    }
    count
}

#[aoc(day2, part2)]
pub fn part2(input: &Vec<Entry>) -> usize {
    let mut count = 0;
    for e in input.iter() {
        let num = e.password.chars()
            .enumerate()
            .filter(|&(i, _)| i+1 == e.min || i+1 == e.max)
            .map(|(_, e)| e)
            .filter(|c| *c == e.letter)
            .count();
        if num == 1 {
            count += 1;
        }
    }
    count
}
