package semantics.exp.red;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.*;
import semantics.value.BoolVal;

import static org.junit.jupiter.api.Assertions.*;

public class LessThanTest {

    @Test
    void create() {
        LessThan lessThan1 = new LessThan(new Num(10), new Num(20));
        LessThan lessThan2 = new LessThan(new Num(10), new Num(20));
        assertEquals(lessThan1, lessThan2);
    }

    @Test
    void create_left_right() {
        LessThan lessThan = new LessThan(new Num(1), new Num(2));
        assertEquals(1, lessThan.getLeft().getReduced().get());
        assertEquals(2, lessThan.getRight().getReduced().get());
    }

    @Test
    void create_reducible() {
        LessThan lessThan = new LessThan(new Num(1), new Num(2));
        assertTrue(lessThan.isReducible());
    }

    @Test
    void create_get_val() {
        LessThan lessThan = new LessThan(new Num(1), new Num(2));
        assertThrows(UnsupportedOperationException.class, lessThan::getReduced);
    }

    @Test
    void reduce_simple_true() {
        LessThan lessThan = new LessThan(new Num(1), new Num(2));
        Expression<BoolVal> reduced = lessThan.reduce(new Environment());
        assertFalse(reduced.isReducible());
        assertTrue(reduced.getReduced().get());
    }

    @Test
    void reduce_simple_false() {
        LessThan lessThan = new LessThan(new Num(5), new Num(1));
        Expression<BoolVal> reduced = lessThan.reduce(new Environment());
        assertFalse(reduced.isReducible());
        assertFalse(reduced.getReduced().get());
    }

    @Test
    void reduce_complex_true() {
        LessThan lessThan = new LessThan(
                new Add(new Num(1), new Num(2)),
                new Multiply(new Num(10), new Num(20)));

        Expression<BoolVal> reduced1 = lessThan.reduce(new Environment());
        assertTrue(reduced1.isReducible());
        Expression<BoolVal> reduced2 = reduced1.reduce(new Environment());
        assertTrue(reduced2.isReducible());
        Expression<BoolVal> reduced3 = reduced2.reduce(new Environment());
        assertFalse(reduced3.isReducible());
        assertTrue(reduced3.getReduced().get());
    }
}
