use regex::Regex;

pub fn day3(data: Vec<String>) -> i32 {
   let re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();
    let mut results = Vec::new();
    for input in data {
        for captures in re.captures_iter(&input) {
            let left = captures.get(1).unwrap().as_str().parse::<i32>().unwrap();
            let right = captures.get(2).unwrap().as_str().parse::<i32>().unwrap();
            results.push((left, right));
        }
    }
    let mut sum = 0;
    for (left, right) in results {
        sum += left * right;
    }
    sum
}



pub fn day3_2(data: Vec<String>) -> i32 {
    let re = Regex::new(r"(mul|do|don't)\((\d{1,3})?,?(\d{1,3})?\)").unwrap();

    let mut results = Vec::new();


    for input in data {
        for captures in re.captures_iter(&input) {
            let func_name = captures.get(1).unwrap().as_str();
            let left = captures.get(2).map(|m| m.as_str().parse::<i32>().ok());
            let right = captures.get(3).map(|m| m.as_str().parse::<i32>().ok());
            
            results.push((func_name.to_string(), left, right));
        }
    }
    let mut active = true;
    let mut sum = 0;
    for (func_name, left, right) in results {
        if func_name == "do" {
            active = true;
        } else if func_name == "don't"  {
            active= false;
        } else {
            if active {
                if let Some(Some(l)) = left {
                    if let Some(Some(r)) = right {
                        sum += l*r;
                    }
                } 
            }
        }
    }
    sum
}