fn main() {
    println!("Hello, world!");

    let s = String::from("hello world");
    let i = first_word(&s);
    println!("index: {}", i);

    let w = first_word2(&s);
    println!("word: {}", w);
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }
    return s.len();
}

fn first_word2(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i]
        }
    }
    &s[..]
}
