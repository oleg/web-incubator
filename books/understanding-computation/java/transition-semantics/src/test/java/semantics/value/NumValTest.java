package semantics.value;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class NumValTest {

    @Test
    void create() {
        NumVal numVal = new NumVal(10);
        assertEquals(10, numVal.get());
    }

    @Test
    void create_big() {
        NumVal numVal = new NumVal(1000);
        assertEquals(1000, numVal.get());
    }
//
//    @Test
//    void reducible() {
//        Num num = new Num(1000);
//        assertFalse(num.isReducible());
//    }

//    @Test
//    void reduce() {
//        Num num = new Num(1000);
//        assertThrows(UnsupportedOperationException.class, () -> num.reduce(new Environment()));
//    }

    @Test
    void to_string() {
        assertEquals("Num(10)", new NumVal(10).toString());
        assertEquals("Num(8)", new NumVal(8).toString());
    }

    @Test
    void equals() {
        assertEquals(new NumVal(10), new NumVal(10));
        assertEquals(new NumVal(5), new NumVal(5));
    }
}