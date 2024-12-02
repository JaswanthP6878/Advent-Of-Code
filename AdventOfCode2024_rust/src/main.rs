mod utils;
mod sol;
use utils::{read_file, read_file_i32};

fn main() {
    let result = match read_file_i32("./inputs/2.txt") {
        Ok(val) => val,
        Err(e) => {
            panic!("{}", e);
        }
    };

    let result = sol::day2::day2(result);
    println!("{:?}", result);
    
}
