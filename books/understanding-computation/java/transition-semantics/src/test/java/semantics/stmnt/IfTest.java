package semantics.stmnt;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.Bool;
import semantics.exp.LessThan;
import semantics.exp.Num;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class IfTest {

    @Test
    void create() {
        If anIf = new If(
                new Bool(true),
                new Assign<>("x", new Num(1)),
                new Assign<>("y", new Num(2)));

        assertTrue(anIf.isReducible());
    }

    @Test
    void bool_true() {
        Assign<NumVal> assignX = new Assign<>("x", new Num(1));
        Assign<NumVal> assignY = new Assign<>("y", new Num(2));
        If anIf = new If(new Bool(true), assignX, assignY);

        StatementResult reduce = anIf.reduce(new Environment());
        assertEquals(assignX, reduce.getStatement());
    }

    @Test
    void bool_false() {
        Assign<NumVal> assignX = new Assign<>("x", new Num(1));
        Assign<NumVal> assignY = new Assign<>("y", new Num(2));
        If anIf = new If(new Bool(false), assignX, assignY);

        StatementResult reduce = anIf.reduce(new Environment());
        assertEquals(assignY, reduce.getStatement());
    }

    @Test
    void complex_condition() {
        Assign<NumVal> assignX = new Assign<>("x", new Num(1));
        Assign<NumVal> assignY = new Assign<>("y", new Num(2));
        If anIf = new If(new LessThan(new Num(10), new Num(12)), assignX, assignY);

        StatementResult reduce = anIf.reduce(new Environment());
        assertEquals(
                new If(new Bool(true), assignX, assignY),
                reduce.getStatement());
    }
}