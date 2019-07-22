use std::str;

const LY_BELOW_NUM: i32 = 10000;
const MAX_TRY: i32 = 50;

fn main() {
    let mut num_ly = 0;
    for x in 0..LY_BELOW_NUM {
        if is_lychrel(x.into()) {
            num_ly += 1;
        }
    }
    println!("Result: {}", num_ly);
    // let (flip, r)  = is_palindrome(3);
    // println!("{}, {}", flip, r);
}

fn is_lychrel(x: f64) -> bool {
    let mut iter_x = x;
    let mut result = true;
    let (flip_x, _) = is_palindromic(x);
    iter_x += flip_x;
    for _ in 0..MAX_TRY {
        let (flip_x, is_palin) = is_palindromic(iter_x);
        if is_palin {
            result = false;
            break
        }
        iter_x += flip_x
    }
    result
}

fn is_palindrome(x: i128) -> (i128, bool) {
    let mut div = 1;
    while x/div >= 10 {
        div *= 10;
    }

    let mut mut_x = x;
    let mut is_palin = true;
    let mut flip_x: i128 = 0;
    //let mut padding = 1;
    while mut_x >= 10 {
        let begin = mut_x/div;
        let end = mut_x%10;
        if begin != end {
            is_palin = false;
        }
        //flip_x += (end *div + begin) * padding;
        mut_x = mut_x - (begin*div);
        mut_x /= 10;
        div /= 100;
        //padding *= 10;
    }
    mut_x=x;
    while mut_x != 0 {
        let r = mut_x%10;
        flip_x = flip_x*10+r;
        mut_x/=10;
    }
    
        //flip_x += mut_x*padding;
    
    (flip_x,is_palin)
}

fn is_palindromic(x: f64) -> (f64, bool) {
    let mut result = true;
    let str_x = x.to_string();
    let byte_x = str_x.as_bytes();
    let mut flip_x = byte_x.to_vec();
    let flip_len = flip_x.len();
    for i in 0..flip_len/2 {
        if flip_x[i] != flip_x[flip_len-i-1] {
            result = false;
        }
        let temp = flip_x[i];
        flip_x[i] = flip_x[flip_len-i-1];
        flip_x[flip_len-i-1] = temp;
    }
    let s = match str::from_utf8(&flip_x) {
        Ok(v) => v,
        Err(e) => panic!("Invalid UTF-8 sequence: {}", e),
    };
    (s.parse::<f64>().unwrap_or(0.0).to_owned(), result)
}