use core::fmt;
use std::str::from_utf8;
use std::vec;

use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{alpha1, char},
    character::complete::{digit1, multispace1, newline, space1, u32},
    multi::{many1, separated_list1},
    sequence::{delimited, preceded},
    IResult,
};

// This solution is heavily inspired by https://github.com/ChristopherBiscardi/advent-of-code/blob/c97bb3bfbe954f500070a0ab95c397e58dda9bf2/2022/rust/day-05/src/lib.rs

#[derive(Debug, Clone, Copy)]
struct Crate(char);

struct Move {
    count: u32,
    from: u32,
    to: u32,
}

// Implement Debug ourselves to have a nicer output
impl fmt::Debug for Move {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        writeln!(f, "Move {} from {} to {}", self.count, self.from, self.to)
    }
}

// Wrap the Vec of Vec in a struct to be able to debug it easily
struct Stacks(Vec<Vec<Crate>>);

// Implement Debug ourselves to have a nicer output
impl fmt::Debug for Stacks {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for (i, stack) in self.0.iter().enumerate() {
            writeln!(f, "Stack {}: {:?}", i, stack)?;
        }
        Ok(())
    }
}

fn main() {
    let file_content = std::fs::read_to_string("input.txt").expect("Could not open file input.txt");
    let top_crates = get_stack_top_crates(&file_content);
    println!("Top creates are {top_crates}");
    let top_crates_part_two = get_stack_top_crates_part_two(&file_content);
    println!("Top creates for part 2 are {top_crates_part_two}");
}

// Part 1 of the puzzle
fn get_stack_top_crates(input: &str) -> String {
    let (input, (mut stacks, moves)) = parse_input(&input).unwrap();
    //dbg!(&stacks);
    for m in moves {
        //dbg!(&m);
        let len = stacks.0[m.from as usize].len();
        let drained: Vec<Crate> = stacks.0[m.from as usize]
            .drain((len - m.count as usize)..)
            .rev()
            .collect();
        for c in drained.iter() {
            stacks.0[m.to as usize].push(c.clone());
        }
        //dbg!(&stacks);
    }
    let result: String = stacks
        .0
        .iter()
        .map(|stack| stack.last().unwrap())
        .map(|c| c.0)
        .collect::<String>();
    result
}

// Part 2 of the puzzle
fn get_stack_top_crates_part_two(input: &str) -> String {
    let (input, (mut stacks, moves)) = parse_input(&input).unwrap();
    //dbg!(&stacks);
    for m in moves {
        //dbg!(&m);
        let len = stacks.0[m.from as usize].len();
        let drained: Vec<Crate> = stacks.0[m.from as usize]
            .drain((len - m.count as usize)..)
            .collect();
        for c in drained.iter() {
            stacks.0[m.to as usize].push(c.clone());
        }
        //dbg!(&stacks);
    }
    let result: String = stacks
        .0
        .iter()
        .map(|stack| stack.last().unwrap())
        .map(|c| c.0)
        .collect::<String>();
    result
}


fn parse_input(input: &str) -> IResult<&str, (Stacks, Vec<Move>)> {
    // Parse the stacks of crates
    let (input, crates_horizontal) = separated_list1(newline, parse_crate_line)(input)?;
    //dbg!(&crates_horizontal);
    // Ignore newline and the numbers below the stacks of creates
    let (input, _) = newline(input)?;
    let (input, _numbers) = many1(preceded(space1, digit1))(input)?;
    let (input, _) = multispace1(input)?; // 1 or more space, tabs, carriage return and line feeds

    // The stacks of crates were parsed horizontally, transform them to have the vertical stacks
    let crates_vertical = rotate_crate_stacks(&crates_horizontal);
    //dbg!(&crates_vertical);

    // Parse the moves
    let (input, moves) = separated_list1(newline, parse_move_line)(input)?;
    //dbg!(&moves);

    Ok((input, (Stacks(crates_vertical), moves)))
}

/// Parse a "crate" line, for instance :
/// [N] [C]    
/// will return
/// [Some(Crate('N',),),Some(Crate('C',),),None,]
fn parse_crate_line(line: &str) -> IResult<&str, Vec<Option<Crate>>> {
    let (input, result) = separated_list1(tag(" "), parse_crate)(line)?;
    Ok((input, result))
}

/// Parse a "crate" or an empty slot, for instance :
/// [N]
/// will return
/// Some(Create('N'))
/// while an empty slot "   " will return None
fn parse_crate(i: &str) -> IResult<&str, Option<Crate>> {
    let (i, c) = alt((tag("   "), delimited(char('['), alpha1, char(']'))))(i)?;

    let result = match c {
        "   " => None,
        value => Some(Crate(value.chars().next().unwrap())),
    };
    Ok((i, result))
}

/// Parse a "move" line, for instance :
/// move 1 from 2 to 1
/// will return
/// Move { count: 1, from: 1, to: 0 }
/// from and to are return as 0-based (converted from 1-based)
fn parse_move_line(input: &str) -> IResult<&str, Move> {
    let (input, _) = tag("move ")(input)?;
    let (input, count) = u32(input)?;
    let (input, _) = tag(" from ")(input)?;
    let (input, from) = u32(input)?;
    let (input, _) = tag(" to ")(input)?;
    let (input, to) = u32(input)?;

    Ok((
        input,
        Move {
            count: count,
            from: from - 1,
            to: to - 1,
        },
    ))
}

/// Given a 2 dimensional vector of Option<Crate>, return a new 2 dimensional vector of Crate
/// rotated by -90°.
fn rotate_crate_stacks(stacks: &Vec<Vec<Option<Crate>>>) -> Vec<Vec<Crate>> {
    let mut rotated: Vec<Vec<Option<Crate>>> = vec![];
    // Initialize the rotated vector with empty vectors,
    // there will be as many stacks as we have of elements in one of the horizontal crates lines
    // (here we take the first line)
    for i in 0..stacks[0].len() {
        rotated.push(vec![]);
    }
    for vec in stacks.iter().rev() {
        for (i, c) in vec.iter().enumerate() {
            rotated[i].push(c.clone());
        }
    }
    // Filter out the None values
    let final_rotated: Vec<Vec<Crate>> = rotated
        .iter()
        .map(|vec| vec.iter().filter_map(|v| *v).collect())
        .collect();
    final_rotated
}

#[cfg(test)]
mod test;
