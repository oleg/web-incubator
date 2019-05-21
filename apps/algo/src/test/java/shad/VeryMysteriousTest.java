package shad;

import org.junit.Test;

import static org.junit.Assert.assertArrayEquals;
import static shad.VeryMysterious.invoke;

public class VeryMysteriousTest {

  @Test
  public void testName() throws Exception {
    assertArrayEquals(new int[]{0, 1, 2, 3, 4},
        invoke(new int[]{0, 1, 2, 3, 4}));
  }

  @Test
  public void testName2() throws Exception {
    assertArrayEquals(new int[]{1, 0, 2, 3, 4},
        invoke(new int[]{1, 0, 2, 3, 4}));
  }

  @Test
  public void testName3() throws Exception {
    assertArrayEquals(new int[]{3, 2, 1, 0, 4},
        invoke(new int[]{3, 2, 1, 0, 4}));
  }

  @Test
  public void testName4() throws Exception {
    assertArrayEquals(new int[]{3, 2, 1, 0, 4},
        invoke(new int[]{3, 2, 1, 0, 4}));
  }

  @Test
  public void testName5() throws Exception {
    assertArrayEquals(new int[]{3, 4, 9, 5, 7, 8, 0, 2, 6, 1},
        invoke("6 9 7 0 1 3 8 4 5 2"));
  }
  //3 4 9 5 7 8 0 2 6 1
}
