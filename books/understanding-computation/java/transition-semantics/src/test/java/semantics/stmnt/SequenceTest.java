package semantics.stmnt;

import org.junit.jupiter.api.Test;
import semantics.env.Environment;
import semantics.exec.Machine;
import semantics.exp.Num;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class SequenceTest {
    Sequence sequence = new Sequence(
            new Assign<>("x", new Num(10)),
            new Assign<>("y", new Num(20)));

    @Test
    void reducible() {
        assertTrue(sequence.isReducible());
    }

//    @Test
//    void reduce() {
//        sequence.reduce(new Environment());
//    }

    @Test
    void execute() {

        Environment result = new Machine().run(sequence, new Environment());

        assertEquals(new NumVal(10), result.get("x"));
        assertEquals(new NumVal(20), result.get("y"));
    }
}