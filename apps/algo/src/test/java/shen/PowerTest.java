package shen;

import org.junit.Before;
import org.junit.Test;

import java.math.BigDecimal;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class PowerTest {

  private Power power;

  @Before
  public void setUp() throws Exception {
    power = new RecursionPower();
  }

  @Test
  public void test_pow() throws Exception {
    assertThat(power.get(m(-1), m(0)), is(m(1)));
    assertThat(power.get(m(0), m(0)), is(m(1)));
    assertThat(power.get(m(1), m(0)), is(m(1)));
    assertThat(power.get(m(2), m(0)), is(m(1)));
    assertThat(power.get(m(100), m(0)), is(m(1)));
  }

  @Test
  public void test_negative_base() throws Exception {
    assertThat(power.get(m(-1), m(1)), is(m(-1)));
    assertThat(power.get(m(-1), m(2)), is(m(1)));
    assertThat(power.get(m(-1), m(3)), is(m(-1)));
  }


  @Test
  public void test_m() throws Exception {
    assertThat(power.get(m(1), m(1)), is(m(1)));
    assertThat(power.get(m(1), m(2)), is(m(1)));
    assertThat(power.get(m(1), m(4)), is(m(1)));
  }

  @Test
  public void test() throws Exception {
    assertThat(power.get(m(2), m(1)), is(m(2)));
    assertThat(power.get(m(2), m(2)), is(m(4)));
    assertThat(power.get(m(2), m(4)), is(m(16)));
  }

  @Test
  public void test_different() throws Exception {
    assertThat(power.get(m(5), m(2)), is(m(25)));
    assertThat(power.get(m(2), m(8)), is(m(256)));
    assertThat(power.get(m(3), m(7)), is(m(2187)));
    assertThat(power.get(m(3), m(8)), is(m(6561)));
    assertThat(power.get(m(4), m(12)), is(m(16777216)));
    assertThat(power.get(m(4), m(13)), is(m(67108864)));
  }

  @Test
  public void benchmark() throws Exception {
    //System.out.println(power.get(m(123), m(994)));
  }


  @Test(expected = IllegalArgumentException.class)
  public void negative_pow() throws Exception {
    power.get(m(2), m(-1));
  }

  private BigDecimal m(long l) {
    return BigDecimal.valueOf(l);
  }
}
