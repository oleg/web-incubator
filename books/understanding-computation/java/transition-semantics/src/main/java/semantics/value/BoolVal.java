package semantics.value;

import lombok.EqualsAndHashCode;

@EqualsAndHashCode
public class BoolVal implements Val<Boolean> {
    private final Boolean value;

    public BoolVal(boolean value) {
        this.value = value;
    }

    @Override
    public Boolean get() {
        return value;
    }

    @Override
    public String toString() {
        return "Bool(" + value + ")";
    }

}
