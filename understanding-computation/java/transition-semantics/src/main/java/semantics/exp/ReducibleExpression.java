package semantics.exp;

import semantics.value.Val;

public abstract class ReducibleExpression<T extends Val<?>> implements Expression<T> {

    @Override
    public boolean isReducible() {
        return true;
    }

    @Override
    public T getReduced() {
        throw new UnsupportedOperationException();
    }
}
