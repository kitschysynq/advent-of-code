use aoc_runner_derive::{aoc, aoc_generator};

#[aoc_generator(day3)]
pub fn generator(input: &str) -> Vec<Vec<bool>> {
    input.lines()
        .map(|l| {
            l.chars()
                .map(|c| if c == '.' { false } else { true })
                .collect()
        })
        .collect()
}

#[aoc(day3, part1)]
pub fn part1(input: &Vec<Vec<bool>>) -> usize {
    let mut count = 0;
    let mut pos = 0;
    for l in input.iter().skip(1) {
        pos = (pos + 3) % l.len();
        count += if l[pos] { 1 } else { 0 };
    }
    count
}

#[aoc(day3, part2)]
pub fn part2(input: &Vec<Vec<bool>>) -> usize {
    let slopes = vec![
        (1usize, 1usize),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2),
    ];

    let mut total = 1;

    for (right, down) in slopes.iter() {
        let mut count = 0;
        let mut pos = 0;
        for l in input.iter().skip(*down).step_by(*down) {
            pos = (pos + *right) % l.len();
            count += if l[pos] { 1 } else { 0 };
        }
        total *= count;
    }
    total
}

#[aoc(day3, part2, take_two)]
pub fn part2_two(input: &Vec<Vec<bool>>) -> usize {
    let slopes = vec![
        (1usize, 1usize),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2),
    ];

    slopes.iter()
        .map(|(right, down)| {
            let mut count = 0;
            let mut pos = 0;
            for l in input.iter().skip(*down).step_by(*down) {
                pos = (pos + *right) % l.len();
                count += if l[pos] { 1 } else { 0 };
            }
            count
        })
        .fold(1, |acc, x| acc * x)
}

