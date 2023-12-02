use std::fs::File;
use std::io::{BufRead, BufReader};
use std::path::PathBuf;

use clap::Parser;
use onig::Regex;

const NUMBER_REGEX: &str = "(?=(zero|one|two|three|four|five|six|seven|eight|nine|\\d))";

#[derive(Parser)]
struct Cli {
    file_path: PathBuf,
}

fn parse_matches(string: &str) -> Option<u32> {
    match string {
        "one" => Some(1),
        "two" => Some(2),
        "three" => Some(3),
        "four" => Some(4),
        "five" => Some(5),
        "six" => Some(6),
        "seven" => Some(7),
        "eight" => Some(8),
        "nine" => Some(9),
        _ if { string.len() == 1 } => string.parse::<u32>().ok(),
        _ => None,
    }
}

fn main() {
    let args = Cli::parse();

    let puzzle_input = File::open(&args.file_path).expect("could not open file");
    let puzzle_input_buffer = BufReader::new(&puzzle_input);

    let re = Regex::new(NUMBER_REGEX).unwrap();

    let num: u32 = puzzle_input_buffer.lines()
        .filter_map(|l| l.ok())
        .map(|line| {
            let test: Vec<u32> = re
                .captures_iter(line.as_str())
                .filter(|cap| !cap.is_empty())
                .map(|cap| {
                    cap.iter_pos()
                        .filter_map(|pos| {
                            let p = pos.unwrap();
                            Some(&line[p.0..p.1])
                        })
                        .collect::<Vec<&str>>()
                })
                .flat_map(|caps| {
                    caps.iter()
                        .filter_map(|m| parse_matches(m))
                        .collect::<Vec<u32>>()
                })
                .collect::<Vec<u32>>();
            test
        })
        .map(|num| {
            let first = num[0];
            let last = num[num.len()-1];
            first * 10 + last
        })
        .sum();

    println!("Calibration result: {}", num)
}
