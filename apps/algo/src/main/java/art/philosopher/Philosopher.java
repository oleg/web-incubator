package art.philosopher;

import java.util.Random;
import java.util.concurrent.TimeUnit;

public class Philosopher implements Runnable {
  private int num;
  Chopstick left;
  Chopstick right;
  private boolean eating;

  public Philosopher() {
  }

  public Philosopher(int num, Chopstick left, Chopstick right) {
    this.num = num;
    this.left = left;
    this.right = right;
  }

  public boolean isEating() {
    return eating;
  }

  public void eat() {
    boolean lt = left.setOwner(this);
    boolean rt = right.setOwner(this);
    eating = lt && rt; //TODO call only one???
  }

  public void stopEating() {
    if (eating) {
      left.removeOwner(this);
      right.removeOwner(this);
      eating = false;
    }
  }

  @Override
  public String toString() {
    return "Philosopher{" +
        "num=" + num
//        + ", left=" + left
//        + ", right=" + right
        + '}';
  }

  public void live() {
    while (true) {
      System.out.println(num + " want eat");
      eat();
      if (isEating()) {
        System.out.println(num + " eating");
        silentSleep(1000 * random.nextInt(10));
        System.out.println(num + " stop");
        stopEating();
      }
      silentSleep(1000);
    }
  }

  @Override
  public void run() {
    live();
  }

  private static Random random = new Random();

  private static void silentSleep(int millis) {
    try {
      TimeUnit.MILLISECONDS.sleep(millis);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
  }

}
