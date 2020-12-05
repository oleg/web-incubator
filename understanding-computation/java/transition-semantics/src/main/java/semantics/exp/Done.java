package semantics.exp;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.env.Environment;
import semantics.value.Val;

@ToString
@EqualsAndHashCode
public class Done<V extends Val<?>> implements Expression<V> {
    private final V value;

    public Done(V value) {
        this.value = value;
    }

    @Override
    public V getReduced() {
        return value;
    }

    @Override
    public boolean isReducible() {
        return false;
    }

    @Override
    public Expression<V> reduce(Environment environment) {
        throw new UnsupportedOperationException();
    }
}
