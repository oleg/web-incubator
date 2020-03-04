package lcode;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class TwoSumLessThanKTest {

    @Test
    void test1() {
        int result = new TwoSumLessThanK().twoSumLessThanK(
                new int[]{34, 23, 1, 24, 75, 33, 54, 8},
                60);
        assertEquals(58, result);
    }

    @Test
    void test2() {
        int result = new TwoSumLessThanK().twoSumLessThanK(
                new int[]{10, 20, 30},
                15);
        assertEquals(-1, result);

    }
    @Test
    void test3() {
        int result = new TwoSumLessThanK().twoSumLessThanK(
                new int[]{254,914,110,900,147,441,209,122,571,942,136,350,160,127,178,839,201,386,462,45,735,467,153,415,875,282,204,534,639,994,284,320,865,468,1,838,275,370,295,574,309,268,415,385,786,62,359,78,854,944},
                200);
        assertEquals(198, result);

    }
}