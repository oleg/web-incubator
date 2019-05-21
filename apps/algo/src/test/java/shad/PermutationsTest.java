package shad;

import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

import static java.util.Arrays.asList;
import static junit.framework.Assert.assertEquals;
import static org.hamcrest.CoreMatchers.allOf;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.junit.matchers.JUnitMatchers.hasItem;

public class PermutationsTest {

  @Test
  public void testName() throws Exception {
    List<List<Integer>> result = Permutations.get(1);
    assertEquals(asList(asList(1)), result);
  }

  @Test
  public void testName2() throws Exception {
    List<List<Integer>> result = Permutations.get(1, 2);

    assertThat(result, hasItem(asList(2, 1)));
    assertThat(result, hasItem(asList(1, 2)));
  }

  @Test
  public void testName3() throws Exception {
    List<List<Integer>> result = Permutations.get(1, 2, 3);

    assertThat(result, allOf(
        hasItem(asList(1, 2, 3)),
        hasItem(asList(2, 1, 3)),
        hasItem(asList(2, 3, 1)),
        hasItem(asList(3, 2, 1)),
        hasItem(asList(3, 1, 2)),
        hasItem(asList(1, 3, 2))));
  }

  @Test
  public void test4() {
    final List<List<Integer>> lists = Permutations.get(1, 2, 3, 4);

    assertThat(lists.size(), is(1 * 2 * 3 * 4));
    assertThat(Permutations.get(1, 2, 3, 4, 5).size(), is(1 * 2 * 3 * 4 * 5));
  }

  @Test
  public void big() throws Exception {
    System.out.println("started:");
    for (int i = 0; i < 100; i++) {
      final long start = System.nanoTime();
      List<List<Integer>> result = Permutations.get(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
      System.out.println(System.nanoTime() - start);
    }
  }

  @Ignore
  @Test
  public void out() throws Exception {
    final List<List<Integer>> lists = Permutations.get(1, 2, 3, 4, 5);
    final int size = lists.size();
    System.out.println(size);

    List<List<Integer>> repeated = new ArrayList<>();
    List<List<Integer>> random = new ArrayList<>();


    for (List<Integer> list : lists) {
      if (list.get(0) == 1 ||
          list.get(1) == 2 ||
          list.get(2) == 3 ||
          list.get(3) == 4 ||
          list.get(4) == 5) {
        repeated.add(list);
      } else {
        random.add(list);
      }
    }
    System.out.println("random: " + random.size());
    for (List<Integer> integers : random) {
      System.out.println(integers);
    }
    System.out.println("repeatable: " + repeated.size());
    for (List<Integer> integers : repeated) {
      System.out.println(integers);
    }

  }
}
