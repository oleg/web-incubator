package sort;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameter;
import org.junit.runners.Parameterized.Parameters;

import static java.util.Arrays.asList;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.*;

@RunWith(Parameterized.class)
public class IntSorterTest {

    @Parameters
    public static Iterable<Object[]> data() {
        return asList(new Object[][]{
            {new int[]{5, 2, 4, 6, 1, 3}, new int[]{1, 2, 3, 4, 5, 6}},
            {new int[]{31, 41, 59, 26, 41, 58}, new int[]{26, 31, 41, 41, 58, 59}},
            {new int[]{0, 1, 2, 3}, new int[]{0, 1, 2, 3}},
            {new int[]{}, new int[]{}},
            {new int[]{7}, new int[]{7}},
            {new int[]{7, -7, 14}, new int[]{-7, 7, 14}},
        });
    }

    @Parameter(0)
    public int[] input;

    @Parameter(1)
    public int[] expected;


    @Test
    public void insertion_sorter_should_sort() throws Exception {
        assertThat(new InsertionSort().sort(input), is(expected));
    }

    @Test
    public void merge_sorter_should_sort() throws Exception {
        assertThat(new InsertionSort().sort(input), is(expected));
    }
}