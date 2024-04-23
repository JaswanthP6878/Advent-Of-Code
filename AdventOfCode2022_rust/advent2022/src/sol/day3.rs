use std::io::BufRead;


pub fn day3(data : &Vec<String>) -> i32 {
    let mut sol = 0;
    for val in data {
        let prio = find_priority(val);
        // println!("{}", prio);
        sol += prio;
    }
    return sol;
}

pub fn day3_2(data: &Vec<String>) -> i32 {
    let mut sol = 0;
    let len = data.len();
    let mut i = 0;
    while i < len {
        let val =  find_badge( data[i].clone(),data[i+1].clone(), data[i+2].clone());
        println!("{}", val);
        sol += val;
        i += 3;
    }
    return sol;
}

fn find_badge(first: String, second: String, third: String) -> i32 {
    println!("{} {} {} ", first,second, third);
    let mut h1 = [0; 52];
    let first_bytes = first.as_bytes();
    let second_bytes = second.as_bytes();
    let third_bytes = third.as_bytes();
    for j in 0..first_bytes.len() {
        if first_bytes[j] >= 97 && first_bytes[j] <= 122 {
            let index: u8 = first_bytes[j] - 97;
            h1[index as usize] += 1;
        } else {
            let index = first_bytes[j] - 65;
            h1[26 + index as usize] += 1;
        }
    }
    for j in 0..second_bytes.len() {
        if second_bytes[j] >= 97 && second_bytes[j] <= 122 {
            let index: u8 = second_bytes[j] - 97;
            h1[index as usize] += 1;
        } else {
            let index = second_bytes[j] - 65;
            h1[26 + index as usize] += 1;
        }
    }
    for j in 0..third_bytes.len() {
        if third_bytes[j] >= 97 && third_bytes[j] <= 122 {
            let index = third_bytes[j] - 97;
            h1[index as usize] += 1;
        } else {
            let index = third_bytes[j] - 65;
            h1[26 + index as usize] += 1;
        }
    }
    println!("{:?}", h1);
    for i in 0..52 {
        if h1[i] == 3 {
            return 1 + i as i32;
        }
    }
    return -1
}



fn find_priority(data: &str) -> i32 {
    let mut h1  = [0; 52];
    let mut h2 = [0; 52];
    let byte_slice = data.as_bytes();
    // println!("{:?}", byte_slice);
    let size = byte_slice.len();

    for i in 0..size/2 {
        if byte_slice[i] >= 97 && byte_slice[i] <= 122 {
            let index = byte_slice[i] - 97;
            h1[index as usize] = 1;
        } else {
            let index = byte_slice[i]- 65;
            h1[26 + index as usize] = 1;
        }
    }
    for j in size/2..size{
        if byte_slice[j] >= 97 && byte_slice[j] <= 122 {
            let index = byte_slice[j] - 97;
            h2[index as usize] = 1;
        } else {
            let index = byte_slice[j] - 65;
            h2[26 + index as usize] = 1;
        }

    }
    for i in 0..52{
        if h1[i] != 0 && h2[i] != 0 && h1[i]  == h2[i] {
            return 1 + i as i32;
        }
    }
    return -1;
}