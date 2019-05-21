package sort;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertArrayEquals;

public class InsertionSortTest {
    private InsertionSort sorter;

    @Before
    public void setup() {
        sorter = new InsertionSort();
    }

    @Test
    public void testSort() {
        int[] input =  new int[] {5, 2, 4, 6, 1, 3};
        int[] sorted = sorter.sort(input);

        assertArrayEquals(new int[] {1, 2, 3, 4, 5, 6}, sorted);
    }

    @Test
    public void testSort2() {
        int[] input =  new int[] {31, 41, 59, 26, 41, 58};
        int[] sorted = sorter.sort(input);

        assertArrayEquals(new int[] {26, 31, 41, 41, 58, 59}, sorted);
    }

    @Test
    public void testSortDesc() {
        int[] input =  new int[] {5, 2, 4, 6, 1, 3};
        int[] sorted = sorter.sortDesc(input);

        assertArrayEquals(new int[] {6, 5, 4, 3, 2, 1}, sorted);
    }

    @Test
    public void testSortDesc2() {
        int[] input =  new int[] {31, 41, 59, 26, 41, 58};
        int[] sorted = sorter.sortDesc(input);

        assertArrayEquals(new int[] {59, 58, 41, 41, 31, 26}, sorted);
    }
}
