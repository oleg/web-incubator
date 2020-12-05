package semantics.exp;

import semantics.value.BoolVal;

public class Bool extends Done<BoolVal> {
    public Bool(boolean value) {
        super(new BoolVal(value));
    }
}
