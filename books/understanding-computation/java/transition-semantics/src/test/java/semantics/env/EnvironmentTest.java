package semantics.env;

import org.junit.jupiter.api.Test;
import semantics.value.NumVal;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

class EnvironmentTest {

    @Test
    void merge() {
        Environment env = new Environment();
        Environment updEnv = env.merge("x", new NumVal(10));

        assertEquals(new NumVal(10), updEnv.get("x"));
    }

    @Test
    void replace() {
        Environment env = new Environment();
        Environment updEnv = env.merge("x", new NumVal(10));
        Environment reUpdEnv = updEnv.merge("x", new NumVal(20));

        assertEquals(new NumVal(20), reUpdEnv.get("x"));
    }

    @Test
    void immutable() {
        Environment env = new Environment();
        Environment updEnv = env.merge("x", new NumVal(10));
        updEnv.merge("x", new NumVal(20));

        assertNull(env.get("x"));
        assertEquals(new NumVal(10), updEnv.get("x"));
    }
}