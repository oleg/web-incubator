fn main() {
    let user1 = User {
        email: String::from("one@example.org"),
        username: String::from("one1"),
        active: true,
        sign_in_count: 1,
    };
    let user2 = build_user(String::from("a"), String::from("b"));
    println!("Hello, world!");
}

fn build_user(email: String, username: String) -> User {
    User {
        email,
        username,
        active: true,
        sign_in_count: 1,
    }
}

struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}
