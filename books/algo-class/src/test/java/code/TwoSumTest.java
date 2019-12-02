package code;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class TwoSumTest {

    @Test
    void sum_is_correct() {
        int[] nums = {2, 7, 11, 15};
        int[] indices = new TwoSum().twoSum_v1(nums, 9);
        assertArrayEquals(new int[]{0, 1}, indices);
    }

    @Test
    void sum2_is_correct() {
        int[] nums = {2, 7, 11, 15};
        int[] indices = new TwoSum().twoSum(nums, 9);
        assertArrayEquals(new int[]{1, 0}, indices);
    }
}