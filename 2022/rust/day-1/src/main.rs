use std::fs;

fn main() {
    println!("Hello, world!");
    let file_content = fs::read_to_string("input.txt").expect("Could not load file");
    let max_calories = get_max_calories(&file_content);
    println!("Max calories: {max_calories}");
    let max_calories_top_three = get_max_calories_of_top_three(&file_content);
    println!("Max calories of top three Elves: {max_calories_top_three}");
}

fn get_max_calories(data: &str) -> String {
    let result = data
        .split("\n\n")
        .map(|elf_load| {
            elf_load
                .lines()
                .map(|item| item.parse::<u32>().unwrap())
                .sum::<u32>()
        })
        .max()
        .unwrap();
    result.to_string()
}

fn get_max_calories_of_top_three(data: &str) -> String {
    let mut result = data
        .split("\n\n")
        .map(|elf_load| {
            elf_load
                .lines()
                .map(|item| item.parse::<u32>().unwrap())
                .sum::<u32>()
        })
        .collect::<Vec<_>>();
    result.sort_by(|a, b| b.cmp(a));
    let sum: u32 = result.iter().take(3).sum();
    sum.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";

    #[test]
    fn should_return_max_calories() {
        let result = get_max_calories(INPUT);
        assert_eq!(result, "24000");
    }

    #[test]
    fn should_return_max_calories_of_top_three_elves() {
        let result = get_max_calories_of_top_three(INPUT);
        assert_eq!(result, "45000");
    }
}
