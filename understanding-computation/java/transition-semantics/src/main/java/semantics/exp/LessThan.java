package semantics.exp;

import lombok.Value;
import semantics.env.Environment;
import semantics.value.BoolVal;
import semantics.value.NumVal;

@Value
public class LessThan extends ReducibleExpression<BoolVal> {
    private final Expression<NumVal> left;
    private final Expression<NumVal> right;

    @Override
    public Expression<BoolVal> reduce(Environment environment) {
        if (left.isReducible())
            return new LessThan(left.reduce(environment), right);
        if (right.isReducible())
            return new LessThan(left, right.reduce(environment));

        Integer vl = left.getReduced().get();
        Integer vr = right.getReduced().get();
        boolean r = vl < vr;
        return new Done<>(new BoolVal(r));
    }
}
