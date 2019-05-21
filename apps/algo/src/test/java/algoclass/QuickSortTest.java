package algoclass;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;
import sort.IntSorter;
import util.PathUtil;

import java.net.URI;
import java.net.URISyntaxException;
import java.nio.file.Paths;

import static algoclass.QuickSortAssignment.PivotStrategy.FINAL;
import static algoclass.QuickSortAssignment.PivotStrategy.FIRST;
import static algoclass.QuickSortAssignment.PivotStrategy.MEDIAN;

public class QuickSortTest extends Assert {
  private IntSorter sorter;

  @Before
  public void setup() {
    sorter = new QuickSortAssignment(MEDIAN);
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
  public void partition_1() throws Exception {
    final int[] input = {2, 1};
    QuickSort.partition(input, 0, 2);
    assertArrayEquals(new int[]{1, 2}, input);
  }

  @Test
  public void partition_2() throws Exception {
    final int[] input = {2, 1, 3, 4, 7, 5, 6};
    QuickSort.partition(input, 0, 2);
    assertArrayEquals(new int[]{1, 2, 3, 4, 7, 5, 6}, input);
  }

  @Test
  public void partition_3() throws Exception {
    final int[] input = {2, 1, 3, 4, 7, 5, 6};
    QuickSort.partition(input, 0, 3);
    assertArrayEquals(new int[]{1, 2, 3, 4, 7, 5, 6}, input);
  }

  @Test
  public void partition_4() throws Exception {
    final int[] input = {4, 1, 3, 2, 7, 5, 6};
    QuickSort.partition(input, 0, 7);
    assertArrayEquals(new int[]{2, 1, 3, 4, 7, 5, 6}, input);
  }

  @Test
  public void partition_5() throws Exception {
    final int[] input = {4, 1, 3, 2, 7, 5, 6};
    QuickSort.partition(input, 4, 7);
    assertArrayEquals(new int[]{4, 1, 3, 2, 6, 5, 7}, input);
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

  @Test
  public void median_pivot() throws Exception {
    int[] ints = PathUtil.readAllInts(Paths.get(resourceUri("QuickSort.txt")));
    final QuickSortAssignment m = new QuickSortAssignment(MEDIAN);
    m.sort(ints);
    assertEquals(138382, m.getComparisonCount());
  }

  @Test
  public void speed_first() throws Exception {
    int[] ints = PathUtil.readAllInts(Paths.get(resourceUri("QuickSort.txt")));
    final QuickSortAssignment f = new QuickSortAssignment(FIRST);
    f.sort(ints);
    assertEquals(162085, f.getComparisonCount());
  }

  @Test
  public void always_final_pivot() throws Exception {
    int[] ints = PathUtil.readAllInts(Paths.get(resourceUri("QuickSort.txt")));

    final QuickSortAssignment l = new QuickSortAssignment(FINAL);
    l.sort(ints);
    assertEquals(164123, l.getComparisonCount());
  }

  private URI resourceUri(String name) throws URISyntaxException {
    return ClassLoader.getSystemResource(name).toURI();
  }

}
