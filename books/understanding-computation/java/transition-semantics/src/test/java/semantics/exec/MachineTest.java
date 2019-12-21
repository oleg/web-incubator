package semantics.exec;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exp.Add;
import semantics.exp.Multiply;
import semantics.exp.Num;
import semantics.exp.Variable;
import semantics.stmnt.Assign;
import semantics.stmnt.Statement;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MachineTest {

    @Test
    void run() {
        Machine machine = new Machine();
        Statement statement = new Assign<>(
                "ff",
                new Add(
                        new Num(10),
                        new Multiply(new Num(2), new Num(3))));

        Environment result = machine.run(statement, new Environment());

        assertEquals(new NumVal(16), result.get("ff"));
    }

    @Test
    void run2() {
        Machine machine = new Machine();
        Statement statement = new Assign<>(
                "ff",
                new Add(
                        new Num(10),
                        new Multiply(new Num(2), new Variable<>("x"))));

        Environment env = new Environment()
                .merge("x", new NumVal(7));

        Environment result = machine.run(statement, env);

        assertEquals(new NumVal(24), result.get("ff"));
    }
}