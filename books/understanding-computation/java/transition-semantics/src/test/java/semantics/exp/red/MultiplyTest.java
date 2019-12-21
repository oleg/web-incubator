package semantics.exp.red;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.Expression;
import semantics.exp.Multiply;
import semantics.exp.Num;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class MultiplyTest {

    @Test
    void create() {
        Multiply multiply = new Multiply(null, null);
        assertTrue(multiply.isReducible());
    }

    @Test
    void create_non_null() {
        Multiply multiply = new Multiply(new Num(1), new Num(7));
        assertTrue(multiply.isReducible());
    }

    @Test
    void reducible() {
        Multiply multiply = new Multiply(new Num(1), new Num(2));
        assertTrue(multiply.isReducible());
    }

    @Test
    void reduce() {
        Multiply multiply = new Multiply(new Num(1), new Num(2));
        assertEquals(new Num(2), multiply.reduce(new Environment()));
    }

    @Test
    void reduce_left() {
        Expression<NumVal> left = new Multiply(new Num(2), new Num(3));
        Expression<NumVal> right = new Num(4);

        Expression<NumVal> reduced = new Multiply(left, right).reduce(new Environment());

        assertEquals(
                new Multiply(new Num(6), new Num(4)),
                reduced);
    }

    @Test
    void reduce_right() {
        Expression<NumVal> left = new Num(3);
        Expression<NumVal> right = new Multiply(new Num(5), new Num(1));

        Expression reduced = new Multiply(left, right).reduce(new Environment());

        assertEquals(
                new Multiply(new Num(3), new Num(5)),
                reduced);
    }

}