package art.philosopher;

import org.junit.Before;
import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class ChopstickTest {

  private Philosopher p1;
  private Philosopher p2;

  @Before
  public void setUp() throws Exception {
    p1 = new Philosopher();
    p2 = new Philosopher();
  }

  @Test
  public void test_take_return_true_first_tame() throws Exception {
    Chopstick c = new Chopstick();


    assertThat(c.setOwner(p1), is(true));
  }

  @Test
  public void test_take_return_true_always_for_the_same_philosopher() throws Exception {
    Chopstick c = new Chopstick();

    assertThat(c.setOwner(p1), is(true));
    assertThat(c.setOwner(p1), is(true));
    assertThat(c.setOwner(p1), is(true));
  }

  @Test
  public void test_take_return_false_for_all_other_philosophers() throws Exception {
    Chopstick c = new Chopstick();

    assertThat(c.setOwner(p1), is(true));
    assertThat(c.setOwner(p2), is(false));
    assertThat(c.setOwner(p2), is(false));
    assertThat(c.setOwner(p1), is(true));
  }

  @Test
  public void test_set_owner_null_should_remove_owner() throws Exception {
    Chopstick c = new Chopstick();

    assertThat(c.setOwner(p1), is(true));
    assertThat(c.removeOwner(p1), is(true));
    assertThat(c.setOwner(p2), is(true));
  }

  @Test
  public void test_set_owner_null_should_remove_owner_only_by_owner() throws Exception {
    Chopstick c = new Chopstick();

    assertThat(c.setOwner(p1), is(true));
    assertThat(c.removeOwner(p2), is(false));
  }

}
