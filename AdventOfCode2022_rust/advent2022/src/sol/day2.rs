
pub fn day2_1(data : &Vec<String>) -> i32{
    let mut sol = 0;
    for val in data {
        let game: Vec<&str> = val.split(' ').collect();
        if game.len() != 2 {
            println!("{:?}", game);
            panic!("game is in wrong format");
        }

        sol += calculate_game_score_part2(game[0], game[1]);
    }
    return sol;
}

// A -> Rock, B -> Paper, C -> Siccsor
// X -> Rock, Y -> Paper, Z -> Siccsor
fn calculate_game_score(opp: &str, mine: &str) -> i32 {
    if mine == "X" { 
        if opp == "A" {
            return 4
        }
        else if opp == "B" {
            return 1
        }
        else {
            return 7
        }
    } else if mine == "Y" {
        if opp == "A" {
            return 8
        }
        else if opp == "B" {
            return 5
        }
        else {
            return 2
        }

    } else {
        if opp == "A" {
            return 3
        }
        else if opp == "B" {
            return 9
        }
        else {
            return 6
        }
    }
}


// part-2
// A -> Rock, B -> Paper, C -> Siccsor
fn calculate_game_score_part2(opp: &str, mine: &str) -> i32 {
    if mine == "X" {  // X means lose
        if opp == "A" {
            return 3
        }
        else if opp == "B" {
            return 1
        }
        else {
            return 2
        }
    } else if mine == "Y" { // Y means draw
        if opp == "A" {
            return 4
        }
        else if opp == "B" {
            return 5
        }
        else {
            return 6
        }

    } else { // Z means win
        if opp == "A" {
            return 8
        }
        else if opp == "B" {
            return 9
        }
        else {
            return 7
        }
    }
}
