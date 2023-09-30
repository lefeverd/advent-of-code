use std::{cmp::Ordering, fs, str::FromStr};

#[derive(PartialEq, Copy, Clone)]
enum Move {
    Rock = 1,
    Paper = 2,
    Scissors = 3,
}

impl FromStr for Move {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "A" | "X" => Ok(Move::Rock),
            "B" | "Y" => Ok(Move::Paper),
            "C" | "Z" => Ok(Move::Scissors),
            _ => Err("Unknown move".to_string()),
        }
    }
}

impl PartialOrd for Move {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        if self == &Move::Scissors && other == &Move::Rock {
            Some(Ordering::Less)
        } else if self == &Move::Rock && other == &Move::Scissors {
            Some(Ordering::Greater)
        } else {
            Some((*self as u8).cmp(&(*other as u8)))
        }
    }
}

fn main() {
    println!("Hello, world!");
    let file_content = fs::read_to_string("input.txt").expect("Could not load file");
    let total = get_total(&file_content);
    println!("Total {total}");
    let total_part2 = get_total_part2(&file_content);
    println!("Total part 2 {total_part2}");
}

fn get_total(data: &str) -> u32 {
    let result = data
        .lines()
        .map(|line| {
            let moves: Vec<Move> = line
                .split(" ")
                .map(|s| s.parse::<Move>().unwrap())
                .collect();
            match moves[0].partial_cmp(&moves[1]) {
                Some(Ordering::Equal) => 3 + moves[1] as u32,
                Some(Ordering::Less) => 6 + moves[1] as u32,
                Some(Ordering::Greater) => 0 + moves[1] as u32,
                None => panic!("moves should be comparable"),
            }
        })
        .sum();
    result
}

fn get_total_part2(data: &str) -> u32 {
    let result = data
        .lines()
        .map(|line| {
            let moves: Vec<&str> = line.split(" ").collect();
            let opponent_move = moves[0].parse::<Move>().unwrap();
            match moves[1] {
                "X" => match opponent_move {
                    Move::Paper => Move::Rock as u32,
                    Move::Rock => Move::Scissors as u32,
                    Move::Scissors => Move::Paper as u32,
                    _ => panic!("Not a valid move"),
                },
                "Y" => 3 + opponent_move as u32,
                "Z" => {
                    6 + match opponent_move {
                        Move::Paper => Move::Scissors as u32,
                        Move::Rock => Move::Paper as u32,
                        Move::Scissors => Move::Rock as u32,
                        _ => panic!("Not a valid move"),
                    }
                }
                _ => panic!("Unexpected order"),
            }
        })
        .sum();
    result
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT: &str = "A Y
B X
C Z";

    #[test]
    fn total_score_should_match() {
        let total = get_total(INPUT);
        assert_eq!(total, 15)
    }

    #[test]
    fn total_score_should_match_part2() {
        let total = get_total_part2(INPUT);
        assert_eq!(total, 12)
    }
}
