package semantics.exp;

import semantics.value.NumVal;

public class Num extends Done<NumVal> {
    public Num(Integer value) {
        super(new NumVal(value));
    }
}
