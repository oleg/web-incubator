package shad;

import org.junit.Test;

import static org.junit.Assert.assertArrayEquals;

public class MsortTest {

  @Test
  public void one() throws Exception {
    assertArrayEquals(new int[]{1}, new Msort().mergesort(new int[]{1}));
  }

  @Test
  public void two() throws Exception {
    assertArrayEquals(new int[]{1, 2}, new Msort().mergesort(new int[]{1, 2}));
    assertArrayEquals(new int[]{1, 2}, new Msort().mergesort(new int[]{2, 1}));
  }

  @Test
  public void a_lot() throws Exception {
    assertArrayEquals(new int[]{2, 3, 4, 5, 6}, new Msort().mergesort(new int[]{3, 4, 2, 5, 6}));
  }

  @Test
  public void from_test() throws Exception {
    assertArrayEquals(new int[]{1, 2, 3, 4, 5, 6}, new Msort().mergesort(new int[]{3, 4, 5, 2, 1, 6}));
  }
}
