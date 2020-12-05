package art.philosopher;

public class Chopstick {

  private int num;

  private Philosopher owner;

  public Chopstick() {
  }

  public Chopstick(int num) {
    this.num = num;
  }

  //TODO check not null
  //TODO sync
  //TODO should be just syncing
  public boolean setOwner(Philosopher philosopher) {
    if (owner != null) {
      return owner == philosopher;//TODO equals
    }
    owner = philosopher;
    return true;
  }

  public boolean removeOwner(Philosopher philosopher) {
    if (owner == null) {
      return true;//or return false;???
    }
    if (owner != philosopher) {
      return false;
    }
    owner = null;
    return true;
  }

  //put should be able only owner!!
  public void put() {
//    taken = false;
  }


  @Override
  public String toString() {
    return "Chopstick{" +
        "num=" + num +
        '}';
  }


}
