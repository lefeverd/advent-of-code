use super::*;

const INPUT: &str = "    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

#[test]
fn test_get_stacks_top_crates() {
    assert_eq!(get_stack_top_crates(INPUT), "CMZ");
}

#[test]
fn test_get_stacks_top_crates_part_two() {
    assert_eq!(get_stack_top_crates_part_two(INPUT), "MCD")
}
