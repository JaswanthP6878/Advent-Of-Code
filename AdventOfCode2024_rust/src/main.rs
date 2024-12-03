mod utils;
mod sol;
use utils::{read_file, read_file_i32};

fn main() {
    let result = match read_file("./inputs/3.txt") {
        Ok(val) => val,
        Err(e) => {
            panic!("{}", e);
        }
    };

    let result = sol::day3::day3_2(result);
    println!("{:?}", result);
    
}
