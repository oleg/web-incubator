package semantics.value;

import lombok.EqualsAndHashCode;
import semantics.exp.Done;
import semantics.exp.Num;

@EqualsAndHashCode
public class NumVal implements Val<Integer> {
    private final Integer value;

    public NumVal(Integer value) {
        this.value = value;
    }

    @Override
    public Integer get() {
        return value;
    }

    @Override
    public String toString() {
        return "Num(" + value + ")";
    }

}
