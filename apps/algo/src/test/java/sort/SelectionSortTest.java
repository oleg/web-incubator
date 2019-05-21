package sort;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertArrayEquals;

public class SelectionSortTest {
    private SelectionSort sorter;

    @Before
    public void setup() {
        sorter = new SelectionSort();
    }

    @Test
    public void testSort() {
        int[] input = {5, 2, 4, 6, 1, 3};
        int[] sorted = sorter.sort(input);

        assertArrayEquals(new int[] {1, 2, 3, 4, 5, 6}, sorted);
    }

    @Test
    public void testSort2() {
        int[] input =  {31, 41, 59, 26, 41, 58};
        int[] sorted = sorter.sort(input);

        assertArrayEquals(new int[] {26, 31, 41, 41, 58, 59}, sorted);
    }
}
