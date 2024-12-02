use std::{fs::File, io::{self, BufRead, BufReader}};

pub fn read_file(path: &str) -> Result<Vec<String>, io::Error> {
    let f = File::open(path)?;
    let buf_reader = BufReader::new(f);
    let mut result = vec![];
    for lines in buf_reader.lines() {
        if let Ok(val) = lines {
            result.push(val);
        }
    }

    Ok(result)
}


// read a Vec<Vec<i32>>
pub fn read_file_i32(path: &str) -> Result<Vec<Vec<i32>>, io::Error> {
    let f = File::open(path)?;
    let buf_reader = BufReader::new(f);
    let mut result: Vec<Vec<i32>> = vec![];
    for lines in buf_reader.lines() {
        if let Ok(val) = lines {
            let mut tmp: Vec<i32> = vec![];
            for i in val.split_whitespace() {
                tmp.push(i.parse().unwrap());
            }
            result.push(tmp);
        }
    }

    Ok(result)
}


#[cfg(test)]
mod tests {
    #[test]
    fn test_readline() {

    }
}