mod utils;
mod sol;
use utils::read_file;

fn main() {
    let result = match read_file("./inputs/1.txt") {
        Ok(val) => val,
        Err(e) => {
            panic!("{}", e);
        }
    };

    let result = sol::day1::day1_2(result);
    println!("{:?}", result);
    
}
