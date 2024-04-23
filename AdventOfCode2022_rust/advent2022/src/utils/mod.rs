use std::fs::File;
use std::io::{BufReader, BufRead};
use std::path::Path;

pub fn read_lines_from_file(path: &Path) -> Vec<String> {
    let mut values = Vec::<String>::new();
    let file = match File::open(path) {
        Ok(file) =>  file,
        _ => panic!()
    };
    let reader = BufReader::new(file);
    for line in reader.lines() {
        if let Ok(val) = line  {
            values.push(val);
        }
    }
    return values
}

