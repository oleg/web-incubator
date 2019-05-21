package art.philosopher;

import org.junit.Ignore;
import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

public class PhilosopherTest {

  @Test
  public void test_thinking_by_default() throws Exception {
    Philosopher phil = new Philosopher(0, null, null);

    assertThat(phil.isEating(), is(false));
  }

  @Test
  public void test_eating_should_call_take() throws Exception {
    Chopstick c0 = mock(Chopstick.class);
    Chopstick c1 = mock(Chopstick.class);

    Philosopher phil = new Philosopher(0, c0, c1);

    phil.eat();

    verify(c0).setOwner(phil);
    verify(c1).setOwner(phil);
  }

  @Test
  public void test_should_start_eating_if_has_two_chopsticks() throws Exception {
    Chopstick c0 = mock(Chopstick.class);
    Philosopher phil = new Philosopher(0, c0, c0);
    when(c0.setOwner(phil)).thenReturn(true);

    phil.eat();

    assertThat(phil.isEating(), is(true));
  }

  @Test
  public void test_should_not_start_eating_if_has_only_one_chopsticks() throws Exception {
    Chopstick c0 = mock(Chopstick.class);
    Chopstick c1 = mock(Chopstick.class);

    Philosopher phil = new Philosopher(0, c0, c1);
    when(c1.setOwner(phil)).thenReturn(true);

    phil.eat();

    assertThat(phil.isEating(), is(false));
  }


  @Test
  public void if_one_take_than_other_can_not_take() throws Exception {
    Chopstick c0 = new Chopstick();
    Chopstick c1 = new Chopstick();

    Philosopher p1 = new Philosopher(0, c0, c1);

    p1.eat();
    assertThat(p1.isEating(), is(true));

    p1.eat();
    assertThat(p1.isEating(), is(true));
  }

  @Test
  public void test_two_philosophers() throws Exception {
    Chopstick c0 = new Chopstick();
    Chopstick c1 = new Chopstick();

    Philosopher p1 = new Philosopher(0, c0, c1);
    Philosopher p2 = new Philosopher(1, c1, c0);

    p1.eat();
    assertThat(p1.isEating(), is(true));
    p1.stopEating();

    p2.eat();
    assertThat(p2.isEating(), is(true));
  }

  @Test
  public void test_stop_eating() throws Exception {
    Chopstick c0 = new Chopstick();
    Chopstick c1 = new Chopstick();

    Philosopher p1 = new Philosopher(0, c0, c1);

    p1.eat();
    assertThat(p1.isEating(), is(true));
    p1.stopEating();
    assertThat(p1.isEating(), is(false));
  }

  @Ignore
  @Test
  public void test_Name() throws Exception {
    Chopstick c0 = new Chopstick(0);
    Chopstick c1 = new Chopstick(1);

    Philosopher p1 = new Philosopher(0, c0, c1);
    Philosopher p2 = new Philosopher(1, c1, c0);

    new Thread(p1).start();
    new Thread(p2).start();

    Thread.currentThread().join();
  }
}
