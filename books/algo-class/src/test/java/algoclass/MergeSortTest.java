package algoclass;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertArrayEquals;

public class MergeSortTest {
  private MergeSort sorter;

  @Before
  public void setup() {
    sorter = new MergeSort();
  }

  @Test
  public void if_empty() throws Exception {
    assertArrayEquals(new int[]{}, sorter.sort(new int[]{}));
  }

  @Test
  public void if_null() throws Exception {
    assertArrayEquals(null, sorter.sort(null));
  }

  @Test
  public void sort_sorted() throws Exception {
    assertArrayEquals(new int[]{1}, sorter.sort(new int[]{1}));
    assertArrayEquals(new int[]{1, 2}, sorter.sort(new int[]{1, 2}));
    assertArrayEquals(new int[]{1, 2, 5, 19}, sorter.sort(new int[]{1, 2, 5, 19}));
  }

  @Test
  public void unsorted_two_elements() throws Exception {
    assertArrayEquals(new int[]{1, 2}, sorter.sort(new int[]{2, 1}));
    assertArrayEquals(new int[]{2, 3}, sorter.sort(new int[]{3, 2}));
  }

  @Test
  public void four_elements() throws Exception {
    assertArrayEquals(new int[]{1, 2, 3, 4}, sorter.sort(new int[]{2, 3, 4, 1}));
  }

  @Test
  public void three_elements() throws Exception {
    assertArrayEquals(new int[]{2, 3, 4}, sorter.sort(new int[]{4, 3, 2}));
  }


  @Test
  public void testSort() {
    int[] input = new int[]{5, 2, 4, 6, 1, 3};
    int[] sorted = sorter.sort(input);

    assertArrayEquals(new int[]{1, 2, 3, 4, 5, 6}, sorted);
  }

  @Test
  public void testSort2() {
    int[] input = new int[]{31, 41, 59, 26, 41, 58, 11};
    int[] sorted = sorter.sort(input);

    assertArrayEquals(new int[]{11, 26, 31, 41, 41, 58, 59}, sorted);
  }


}
