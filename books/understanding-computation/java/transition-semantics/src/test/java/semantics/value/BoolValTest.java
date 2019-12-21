package semantics.value;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

class BoolValTest {

    @Test
    void create_false() {
        BoolVal boolVal = new BoolVal(false);
        assertFalse(boolVal.get());
    }

    @Test
    void create_true() {
        BoolVal boolVal = new BoolVal(true);
        assertTrue(boolVal.get());
    }
}