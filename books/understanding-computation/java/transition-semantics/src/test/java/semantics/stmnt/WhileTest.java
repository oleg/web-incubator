package semantics.stmnt;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exec.Machine;
import semantics.exp.Add;
import semantics.exp.LessThan;
import semantics.exp.Num;
import semantics.exp.Variable;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.*;

class WhileTest {

    While aWhile = new While(
            new LessThan(new Variable<>("x"), new Num(20)),
            new Assign<>("x", new Add(new Variable<>("x"), new Num(1))));

    @Test
    void reducible() {
        assertTrue(aWhile.isReducible());
    }

    @Test
    void reduce() {
        StatementResult reduced = aWhile.reduce(new Environment().merge("x", new NumVal(18)));
        assertNotNull(reduced.getStatement());
        assertNotNull(reduced.getEnvironment());
    }

    @Test
    void reduce_result() {
        Environment env = new Environment().merge("x", new NumVal(18));

        Environment resultEnv = new Machine().run(aWhile, env);

        assertEquals(
                new Environment().merge("x", new NumVal(20)),
                resultEnv);
    }


}