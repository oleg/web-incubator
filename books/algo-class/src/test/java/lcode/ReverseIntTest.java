package lcode;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ReverseIntTest {

    @Test
    void test_positive() {
        assertEquals(321, new ReverseInt().reverse(123));
    }

    @Test
    void test_negative() {
        assertEquals(-321, new ReverseInt().reverse(-123));
    }

    @Test
    void test_zero() {
        assertEquals(21, new ReverseInt().reverse(120));
    }

    @Test
    void test_big() {
        assertEquals(0, new ReverseInt().reverse(1534236469));
    }

    @Test
    void name() {
        //int j = (n - 1) >> 1;
        System.out.println("len   n n-1   j   k");
        for (int len = 0; len < 11; len++) {
            int n = len - 1;
            int j = (n - 1) >> 1;
            int k = n - j;
            System.out.printf("%3d %3d %3d %3d %3d%n", len, n, (n - 1), j, k);
        }
    }

    @Test
    void name1() {
        int[] val = new int[]{1, 2, 3, 4};
        int n = val.length - 1;
        for (int j = (n - 1) >> 1; j >= 0; j--) {
            int k = n - j;
            int cj = val[j];
            val[j] = val[k];
            val[k] = cj;
        }

    }
}