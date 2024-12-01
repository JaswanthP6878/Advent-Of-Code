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


#[cfg(test)]
mod tests {
    #[test]
    fn test_readline() {

    }
}