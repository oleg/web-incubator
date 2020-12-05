package semantics.exp;

import lombok.Value;
import semantics.env.Environment;
import semantics.value.NumVal;

@Value
public class Multiply extends ReducibleExpression<NumVal> {

    private final Expression<NumVal> left;
    private final Expression<NumVal> right;

    @Override
    public Expression<NumVal> reduce(Environment environment) {
        if (left.isReducible())
            return new Multiply(left.reduce(environment), right);

        if (right.isReducible())
            return new Multiply(left, right.reduce(environment));

        Integer vl = left.getReduced().get();
        Integer vr = right.getReduced().get();
        Integer r = vl * vr;
        return new Done<>(new NumVal(r));
    }

}
