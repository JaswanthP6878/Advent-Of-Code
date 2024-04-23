
pub fn main(data: &Vec<String>) -> i32 {
    let mut max_count = 0;
    let mut tmp_count = 0;
    for val in data {
        // println!("{}, length: {}", *val, val.len());
        if val.len() == 0 {
            if tmp_count > max_count {
                max_count = tmp_count;
            }
            tmp_count = 0;
        } else {
            tmp_count += val.parse::<i32>().expect("cannot prase this stirng");
        }
    }
    if tmp_count > max_count {
        max_count = tmp_count;
    }
    return max_count;

}



fn create_calories(data: &Vec<String>) -> Vec<i32> {
    let mut sol = Vec::new();
    let mut tmp_count = 0;
    for val in data {
        if val.len() == 0 {
            sol.push(tmp_count);
            tmp_count = 0;
        } else {
            tmp_count += val.parse::<i32>().expect("cannot prase this stirng");
        }
    }
    sol.push(tmp_count);
    return sol;
}


pub fn day1_2(data: &Vec<String>) -> i32 {
    let mut cals_per_elve: Vec<i32> = create_calories(data);
    println!("{:?}", cals_per_elve);
    cals_per_elve.sort();
    let length = cals_per_elve.len();

    return cals_per_elve[length-1] + cals_per_elve[length - 2 ] + cals_per_elve[length - 3];

}