use super::*;

const INPUT: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

#[test]
fn test_fully_contains() {
    assert_eq!(fully_contains_count(INPUT), 2);
}

#[test]
fn test_overlaps() {
    assert_eq!(overlaps(INPUT), 4);
}
