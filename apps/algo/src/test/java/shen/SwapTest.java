package shen;

import org.junit.Before;
import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class SwapTest {

  private int a;
  private int b;

  @Before
  public void setUp() throws Exception {
    a = 10;
    b = 20;
  }

  @Test
  public void with_temp_variable() throws Exception {

    int t = a;
    a = b;
    b = t;

    assertThat(a, is(20));
    assertThat(b, is(10));
  }

  @Test
  public void without_temp_variable() throws Exception {
    a = a + b; // a = a + b
    b = a - b; // b = (a + b) - b = a
    a = a - b; // a = (a + b) - a = b

    assertThat(a, is(20));
    assertThat(b, is(10));
  }
}
