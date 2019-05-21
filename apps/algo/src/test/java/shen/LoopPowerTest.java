package shen;

import org.junit.Before;
import org.junit.Test;

import java.math.BigDecimal;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class LoopPowerTest {

  private Power power;

  @Before
  public void setUp() throws Exception {
    power = new LoopPower();
  }

  @Test
  public void test_pow() throws Exception {
    assertThat(power.get(m(-1), m(0)), is(m(1)));
    assertThat(power.get(m(0), m(0)), is(m(1)));
    assertThat(power.get(m(1), m(0)), is(m(1)));
    assertThat(power.get(m(2), m(0)), is(m(1)));
    assertThat(power.get(m(100), m(0)), is(m(1)));
  }

  @Test(expected = IllegalArgumentException.class)
  public void test_negative() throws Exception {
    power.get(m(10), m(-2));
  }

  @Test
  public void test_powers() throws Exception {
    assertThat(power.get(m(2), m(1)), is(m(2)));
    assertThat(power.get(m(44), m(1)), is(m(44)));
    assertThat(power.get(m(-344), m(1)), is(m(-344)));
  }


  private BigDecimal m(long i) {
    return BigDecimal.valueOf(i);
  }
}
