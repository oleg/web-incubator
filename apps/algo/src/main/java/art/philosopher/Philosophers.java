package art.philosopher;

// 5 philosophers
// 5 chairs
// 5 chopsticks

// philosopher think
// than get angry
// get two chopsticks
// eat
// than start think again

// each philosopher is a thread
// each chopstick is a shared object


public class Philosophers {

  public static void main(String... args) throws InterruptedException {
    int max = 5;

    Chopstick[] chopsticks = new Chopstick[max];
    for (int i = 0; i < max; i++) {
      chopsticks[i] = new Chopstick(i);
    }

    Philosopher[] philosophers = new Philosopher[max];
    for (int i = 0; i < max; i++) {
      int left = i;
      int right = i + 1;

      if (right == max) {
        right = 0;
      }

      philosophers[i] = new Philosopher(i, chopsticks[left], chopsticks[right]);
    }

    for (Philosopher philosopher : philosophers) {
//      philosopher.start();
      Thread.sleep(100);
    }

  }

}

