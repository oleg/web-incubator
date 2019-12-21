package semantics.stmnt;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.Add;
import semantics.exp.Num;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.*;

class AssignTest {

    @Test
    void is_reducible() {
        Assign<NumVal> a = new Assign<>("x", new Num(19));
        assertTrue(a.isReducible());
    }

    @Test
    void reduce_updates_environment() {
        Assign<NumVal> a = new Assign<>("x", new Num(19));

        StatementResult result = a.reduce(new Environment());
        Environment newEnv = result.getEnvironment();

        assertEquals(new NumVal(19), newEnv.get("x"));
    }

    @Test
    void reduce_returns_do_nothing() {
        Assign<NumVal> a = new Assign<>("x", new Num(19));

        StatementResult result = a.reduce(new Environment());
        Statement statement = result.getStatement();

        assertFalse(statement.isReducible());
        assertEquals(DoNothing.INSTANCE, statement);
    }

    @Test
    void reduce_reduces_expression() {
        Assign<NumVal> a = new Assign<>("x", new Add(new Num(5), new Num(50)));
        Environment env = new Environment();

        StatementResult result = a.reduce(env);
        assertEquals(env, result.getEnvironment());
        assertEquals(new Assign<>("x", new Num(55)), result.getStatement());
    }

}