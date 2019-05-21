function* fibonacci() {
    let prev = 0, next = 1;

    for (;;) {
        yield next;
        [next, prev] = [next + prev, next];
    }
}


let x = fibonacci()
for (let y = 0; y < 10000; y++)
    console.log(x.next());