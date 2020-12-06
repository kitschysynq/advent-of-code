use aoc_runner_derive::aoc_lib;

pub mod day1;
pub mod day2;
pub mod day3;
    
aoc_lib!{ year = 2020 }

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
