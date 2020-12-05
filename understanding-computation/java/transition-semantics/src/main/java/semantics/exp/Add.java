package semantics.exp;

import lombok.Value;
import semantics.env.Environment;
import semantics.value.NumVal;

@Value
public class Add extends ReducibleExpression<NumVal> {

    private final Expression<NumVal> left;
    private final Expression<NumVal> right;

    @Override
    public Expression<NumVal> reduce(Environment environment) {
        if (left.isReducible()) {
            return new Add(left.reduce(environment), right);
        }
        if (right.isReducible()) {
            return new Add(left, right.reduce(environment));
        }
        Integer vl = left.getReduced().get();
        Integer vr = right.getReduced().get();
        int r = vl + vr;
        return new Done<>(new NumVal(r));
    }

}
