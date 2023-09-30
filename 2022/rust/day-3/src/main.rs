use std::fs;

// A more elegant solution can be found at https://nickymeuleman.netlify.app/garden/aoc2022-day03

fn main() {
    println!("Hello, world!");
    let file_content: String =
        fs::read_to_string("input.txt").expect("Could not read file input.txt");
    let sum: u32 = get_priorities_sum(&file_content);
    println!("Sum is {sum}");
    let sum_badges: u32 = get_badges(&file_content);
    println!("Sum for badges is {sum_badges}");
}

fn get_priorities_sum(input: &str) -> u32 {
    let sum = input
        .lines()
        .map(|line| {
            let len: usize = line.chars().count();
            let middle = len / 2;
            let first: &str = &line[0..middle];
            let second: &str = &line[middle..];
            for letter in first.chars() {
                if second.contains(letter) {
                    return get_char_position(&letter) as u32;
                }
            }
            1
        })
        .sum();
    return sum;
}

fn get_char_position(letter: &char) -> u16 {
    // char a is 97, and we should return 1
    let mut offset = 96;
    if letter.is_uppercase() {
        // char A is 65
        // and we should return 27
        offset = 64 - 26;
    }
    (*letter as u16) - offset
}

fn get_badges(input: &str) -> u32 {
    let sum: u32 = input
        .lines()
        .map(|line| line.as_bytes())
        .collect::<Vec<_>>()
        .chunks(3)
        .map(|chunks| {
            chunks[0]
                .iter()
                .find(|item| chunks[1].contains(item) && chunks[2].contains(item))
                .unwrap()
        })
        .map(|item| match item {
            b'a'..=b'z' => (item - b'a') as u32 + 1,
            _ => (item - b'A') as u32 + 1 + 26,
        })
        .sum();
    return sum;
}

#[cfg(test)]
mod test;
