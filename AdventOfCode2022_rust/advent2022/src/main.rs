mod utils;
mod sol;

use std::path::Path;
use utils::read_lines_from_file;

fn main() {
    let p = Path::new("./resources/day3_test.txt");
    let data = read_lines_from_file(p);
    // println!("{:?}", &data);
    println!("{}", sol::day3::day3_2(&data));
}
