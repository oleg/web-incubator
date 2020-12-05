package semantics.exp.red;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.Add;
import semantics.exp.Expression;
import semantics.exp.Num;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class AddTest {

    @Test
    void create() {
        Add add = new Add(null, null);
        assertTrue(add.isReducible());
    }

    @Test
    void create_non_null() {
        Add add = new Add(new Num(10), new Num(7));
        assertTrue(add.isReducible());
    }

    @Test
    void reducible() {
        Add add = new Add(new Num(1), new Num(2));
        assertTrue(add.isReducible());
    }

    @Test
    void reduce() {
        Add add = new Add(new Num(1), new Num(2));
        assertEquals(new Num(3), add.reduce(new Environment()));
    }

    @Test
    void reduce_left() {
        Expression<NumVal> left = new Add(new Num(1), new Num(2));
        Expression<NumVal> right = new Num(4);

        Expression<NumVal> reduced = new Add(left, right).reduce(new Environment());

        assertEquals(
                new Add(new Num(3), new Num(4)),
                reduced);
    }

    @Test
    void reduce_right() {
        Expression<NumVal> left = new Num(3);
        Expression<NumVal> right = new Add(new Num(5), new Num(10));

        Expression<NumVal> reduced = new Add(left, right).reduce(new Environment());

        assertEquals(
                new Add(new Num(3), new Num(15)),
                reduced);
    }
}