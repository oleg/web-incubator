package lcode;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class JewelsTest {

    @Test
    void test1() {
        assertEquals(3, new Jewels().numJewelsInStones("aA", "aAAbbbb"));
    }

}
