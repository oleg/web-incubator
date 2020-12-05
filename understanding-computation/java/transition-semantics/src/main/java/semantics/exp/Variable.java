package semantics.exp;

import lombok.Value;
import semantics.env.Environment;
import semantics.value.Val;

@Value
public class Variable<T extends Val<?>> extends ReducibleExpression<T> {
    private final String name;

    @Override
    public Expression<T> reduce(Environment environment) {
        return new Done<>((T) environment.get(name));
    }

}
