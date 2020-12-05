package semantics.exp.red;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.Expression;
import semantics.exp.Variable;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;

class VariableTest {

    @Test
    void reduce_not_found() {
        Variable<NumVal> x = new Variable<>("x");
        Expression<NumVal> reduce = x.reduce(new Environment());
        assertFalse(reduce.isReducible());
    }

    @Test
    void reduce_found() {
        Variable<NumVal> x = new Variable<>("x");
        Expression<NumVal> reduce = x.reduce(new Environment().merge("x", new NumVal(1000)));
        assertFalse(reduce.isReducible());
        assertEquals(new NumVal(1000), reduce.getReduced());
    }
}