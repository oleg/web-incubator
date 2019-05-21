package shad;

import org.junit.Test;

import java.math.BigInteger;
import java.util.List;

import static java.util.Arrays.asList;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.junit.matchers.JUnitMatchers.hasItem;
import static shad.NetPath.Direction.DOWN;
import static shad.NetPath.Direction.LEFT;

public class NetPathTest {
  @Test
  public void directions_2() throws Exception {
    final List<List<NetPath.Direction>> directions = new NetPath(1, 1).directions();

    assertThat(directions.size(), is(2));
    assertThat(directions, hasItem(asList(LEFT, DOWN)));
    assertThat(directions, hasItem(asList(DOWN, LEFT)));
  }

  @Test
  public void directions_3() throws Exception {
    final List<List<NetPath.Direction>> directions = new NetPath(1, 2).directions();

    assertThat(directions.size(), is(3));
    assertThat(directions, hasItem(asList(LEFT, DOWN, DOWN)));
    assertThat(directions, hasItem(asList(DOWN, LEFT, DOWN)));
    assertThat(directions, hasItem(asList(DOWN, DOWN, LEFT)));
  }

  @Test
  public void count() throws Exception {
    assertThat(new NetPath(1, 1).count(), is(2));
    assertThat(new NetPath(1, 2).count(), is(3));
    assertThat(new NetPath(2, 1).count(), is(3));
    assertThat(new NetPath(2, 2).count(), is(6));
    assertThat(new NetPath(3, 2).count(), is(10));
    assertThat(new NetPath(3, 3).count(), is(20));
    assertThat(new NetPath(3, 4).count(), is(35));
    assertThat(new NetPath(4, 4).count(), is(70));
  }

  @Test
  public void count_by_formula() throws Exception {
    assertThat(new NetPath(1, 1).count(), is(netFactI(1, 1)));
    assertThat(new NetPath(1, 2).count(), is(netFactI(1, 2)));
    assertThat(new NetPath(2, 2).count(), is(netFactI(2, 2)));
    assertThat(new NetPath(3, 1).count(), is(netFactI(3, 1)));
    assertThat(new NetPath(3, 2).count(), is(netFactI(3, 2)));
    assertThat(new NetPath(3, 3).count(), is(netFactI(3, 3)));
    assertThat(new NetPath(3, 4).count(), is(netFactI(3, 4)));
    assertThat(new NetPath(4, 4).count(), is(netFactI(4, 4)));
  }

  @Test
  public void asdf() {
    System.out.println(netFact(20, 19));
  }

  private static BigInteger netFact(int a, int b) {
    return factorial(a + b).divide(factorial(a).multiply(factorial(b)));
  }

  private static int netFactI(int a, int b) {
    return netFact(a, b).intValue();
  }

  ///an ad-hoc realisation
  private static BigInteger factorial(int n) {
    BigInteger result = BigInteger.ONE;
    for (long i = 1; i <= n; i++) {
      result = result.multiply(BigInteger.valueOf(i));
    }
    return result;
  }

  @Test
  public void big_count() throws Exception {
    assertThat(new NetPath(5, 5).count(), is(252));
  }
}
