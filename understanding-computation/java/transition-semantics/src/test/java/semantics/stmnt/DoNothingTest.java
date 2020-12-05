package semantics.stmnt;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;

class DoNothingTest {

    @Test
    void reducible() {
        assertFalse(DoNothing.INSTANCE.isReducible());
    }

    @Test
    void equal() {
        assertEquals(DoNothing.INSTANCE, DoNothing.INSTANCE);
    }

}