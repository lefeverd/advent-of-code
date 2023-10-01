use std::{fs, str::FromStr};

// Another interesting method to solve this is by using RangeInclusive : https://fasterthanli.me/series/advent-of-code-2022/part-4

struct Assignment {
    sectionStart: u32,
    sectionEnd: u32,
}

impl FromStr for Assignment {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: Vec<_> = s.split("-").collect();
        Ok(Assignment {
            sectionStart: parts[0].parse().unwrap(),
            sectionEnd: parts[1].parse().unwrap(),
        })
    }
}

impl Assignment {
    fn contains(&self, other: &Assignment) -> bool {
        self.sectionStart >= other.sectionStart && self.sectionEnd <= other.sectionEnd
    }

    fn overlaps(&self, other: &Assignment) -> bool {
        (self.sectionEnd >= other.sectionStart && self.sectionEnd <= other.sectionEnd)
            || (self.sectionStart >= other.sectionStart && self.sectionStart <= other.sectionEnd)
    }
}

fn main() {
    let file_content = fs::read_to_string("input.txt").expect("Cannot read file input.txt");
    let sum = fully_contains_count(&file_content);
    println!("Sum of pairs where one fully contains the other: {sum}");
    let sum_overlaps = overlaps(&file_content);
    println!("Sum of pairs where one overlaps the other: {sum_overlaps}");
}

fn fully_contains_count(input: &str) -> u32 {
    let result: u32 = input
        .lines()
        .map(|line| {
            let assignments: Vec<Assignment> = line
                .split(",")
                .map(|assignment| assignment.parse::<Assignment>().unwrap())
                .collect();
            if assignments[0].contains(&assignments[1]) || assignments[1].contains(&assignments[0])
            {
                return 1;
            }
            0
        })
        .sum();
    result
}

fn overlaps(input: &str) -> u32 {
    let result: u32 = input
        .lines()
        .map(|line| {
            let assignments: Vec<Assignment> = line
                .split(",")
                .map(|assignment| assignment.parse::<Assignment>().unwrap())
                .collect();
            if assignments[0].overlaps(&assignments[1]) || assignments[1].overlaps(&assignments[0])
            {
                return 1;
            }
            0
        })
        .sum();
    result
}

#[cfg(test)]
mod test;
