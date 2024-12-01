use std::collections::HashMap;

pub fn day1(data: Vec<String>) -> i32 { 
    let mut vec1: Vec<i32> = vec![];
    let mut vec2: Vec<i32> = vec![];

    for i in data {
        let mut i_ter = i.split_whitespace();
        if let Some(val) = i_ter.next() {
            vec1.push(val.parse().unwrap());
        }
        if let Some(val) = i_ter.next() {
            vec2.push(val.parse().unwrap());
        }
    }
    // sorting the lists
    vec1.sort();
    vec2.sort();
    assert_eq!(vec1.len(), vec2.len());
    let mut sum = 0;
    for i in 0..vec1.len() {
        let mut diff =  vec1[i] - vec2[i];
        if diff < 0 {
            diff = -1 * diff;
        }
        sum += diff;
    }
    sum

}

pub fn day1_2(data: Vec<String>) -> i32 {
    let mut left_list: Vec<i32> = vec![];
    // create count of right list
    let mut map: HashMap<i32, i32> = HashMap::new();
    for i in data {
        let mut i_ter = i.split_whitespace();
        if let Some(val) = i_ter.next() {
            left_list.push(val.parse().unwrap());
        }
        if let Some(val) = i_ter.next() {
            let val_int: i32 = val.parse().unwrap();
            if let Some(value) = map.get(&val_int) {
                map.insert(val_int, value+1);
            } else {
                map.insert(val_int, 1);
            }

        }
    }
    let mut sum = 0;
    for i in left_list.iter() {
        if let Some(val) = map.get(i) {
            sum += i * val;
        }
    }

    sum
}