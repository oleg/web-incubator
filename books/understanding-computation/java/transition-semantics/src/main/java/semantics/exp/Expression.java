package semantics.exp;

import semantics.env.Environment;
import semantics.value.Val;

public interface Expression<V extends Val<?>> {

    boolean isReducible();

    Expression<V> reduce(Environment environment);

    V getReduced();

}


