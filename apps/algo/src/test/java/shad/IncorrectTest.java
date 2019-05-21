package shad;

import org.junit.Ignore;
import org.junit.Test;

import java.util.List;

import static org.junit.Assert.assertArrayEquals;

public class IncorrectTest {

  @Test
  public void testName() throws Exception {
    final int[] array = {0};
    Incorrect.invoke(array);
    assertArrayEquals(new int[]{0}, array);
  }

  @Test
  public void testName_1() throws Exception {
    final int[] array = {1, 0};
    Incorrect.invoke(array);
    assertArrayEquals(new int[]{0, 1}, array);
  }

  @Test
  public void testName2() throws Exception {
    final int[] array = {4, 3, 2, 1, 0};
    Incorrect.invoke(array);
    assertArrayEquals(new int[]{0, 1, 2, 3, 4}, array);
  }

  @Test
  @Ignore
  public void test_perm() throws Exception {
    final List<List<Integer>> lists = Permutations.get(0, 1, 2, 3, 4);
    for (List<Integer> list : lists) {
      System.out.println(list);
      Incorrect.invoke(list.toArray(new Integer[list.size()]));
    }
  }

  @Test
  public void test_perm_x() throws Exception {
    Incorrect.invoke(new Integer[]{2, 4, 3, 1, 0});
    System.out.println("ok");
  }
}
