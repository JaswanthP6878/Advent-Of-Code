
enum Status {
    Increasing,
    Decreasing,
    NotSafe
}
// how many of the reports are safe;
pub fn day2(data: Vec<Vec<i32>>) -> i32 {
    let mut safe_count = 0;
    for datum in data {
        if is_safe_with_dampner_inneficent(&datum) {
            safe_count += 1;
        }
    }
    safe_count
}

fn is_safe(data: &[i32]) -> bool {
    assert!(data.len() >= 2);
    // condition to check if it should be increasing or decreasing
    let mut is_increasing: bool = false;
    if data[0] < data[1] {
        is_increasing = true;
    }
    if is_increasing {
        for i in 0..data.len()-1 {
            if (data[i] < data[i+1]) && (data[i+1] - data[i] >= 1 && data[i+1] - data[i] <= 3) {
                continue;
            } else {
                return false;
            }
        }
    } else { //decreasing
        for i in 0..data.len()-1 {
            if (data[i] > data[i+1]) && (data[i] - data[i+1] >= 1 && data[i] - data[i+1] <= 3) {
                continue;
            } else {
                return false;
            }
        }
    }
    return true;
}

fn is_safe_with_dampner_inneficent(data: &Vec<i32>) -> bool {
    for i in 0..data.len() {
        let mut permutation = data.clone();
        permutation.remove(i);
        if is_safe(&permutation) {
            return true;
        }
    }
    return false;
}

// allows one level of error
fn is_safe_with_dampner(data: &[i32]) -> bool {
    let is_increasing = check_increasing(&data);
    let mut unsafe_count = 0;
    let mut  unsafe_index: usize = data.len();
    if let Status::Increasing = is_increasing {
        for i in 0..data.len()-1 {
            if data[i] < data[i+1] && valid_gap(data[i], data[i+1]) {
                continue;
            } else {
                unsafe_count += 1;
                unsafe_index = i;
            }
        }

        if unsafe_count == 1 && unsafe_index < data.len() {
            if unsafe_index > 0 {
                if data[unsafe_index - 1] < data[unsafe_index + 1] && valid_gap(data[unsafe_index-1], data[unsafe_index + 1]) {
                    return true;
                } else {
                    return false;
                }
            }
        }
    } else if let Status::Decreasing = is_increasing {
        for i in 0..data.len()-1 {
            if data[i] > data[i+1] && valid_gap(data[i], data[i+1]) {
                continue;
            } else {
                unsafe_count += 1;
                unsafe_index = i;
            }
        }
        if unsafe_count == 1 && unsafe_index < data.len() {
            if unsafe_index > 0 {
                if data[unsafe_index - 1] > data[unsafe_index + 1] && valid_gap(data[unsafe_index-1], data[unsafe_index + 1]) {
                    return true;
                } else {
                    return false;
                }
            }
        }
 
    }
    return false;
}



fn check_increasing(data: &[i32]) -> Status {
    let mut increase_count = 0;
    let mut decrease_count = 0;
    for i in 0..data.len()-1 {
        if data[i] > data[i+1] && valid_gap(data[i], data[i+1]) {
            decrease_count += 1;
        } else if data[i] < data[i+1] && valid_gap(data[i], data[i+1]) {
            increase_count += 1;
        } else {
            continue;
        }
    }
    if increase_count == decrease_count {
        Status::NotSafe
    } else if increase_count < decrease_count {
        Status::Decreasing
    } else {
        Status::Increasing
    }
}

fn valid_gap(d1: i32, d2: i32) -> bool{
    if d1 > d2 {
        return d1 - d2 >= 1 && d1 - d2 <= 3;
    } else {
        return d2 - d1 >= 1 && d2 - d1 <= 3;
    }
}


#[cfg(test)]
mod tests {
    use crate::sol::day2::{is_safe, is_safe_with_dampner};
    #[test]
    fn test_is_safe() {
        let input = vec![7,6,4,2,1];
        println!("{:?}", input);
        assert_eq!(is_safe(&input), true);
    }

    #[test]
    fn test_is_not_safe() {
        let input = vec![1,2,7,8,9];
        println!("{:?}", input);
        assert_eq!(is_safe(&input), false);
    }


    #[test]
    fn test_is_unsafe_with_dampner() {
        let input = vec![1,2,7,8,9];
        println!("{:?}", input);
        assert_eq!(is_safe_with_dampner(&input), false);
    }

    #[test]
    fn test_is_safe_with_dampner() {
        let input = vec![1,3,2,4,5];
        println!("{:?}", input);
        assert_eq!(is_safe_with_dampner(&input), true);
    }
    #[test]
    fn test_is_safe_with_dampner_2() {
        let input = vec![8,6,4,4,1];
        println!("{:?}", input);
        assert_eq!(is_safe_with_dampner(&input), true);
    }
}